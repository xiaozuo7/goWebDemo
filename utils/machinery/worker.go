package machinery

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	"github.com/RichardKnop/machinery/v2/config"
	locksiface "github.com/RichardKnop/machinery/v2/locks/iface"
	"github.com/RichardKnop/machinery/v2/tasks"
	"goWebDemo/utils"
	"log"
)

var cnf = &config.Config{
	Broker:        utils.Broker,
	DefaultQueue:  "machinery_tasks",
	ResultBackend: utils.ResultBackend,
	Lock:          utils.Lock,
	AMQP: &config.AMQPConfig{
		Exchange:     "machinery_exchange",
		ExchangeType: "direct",
		BindingKey:   "machinery_task",
	},
	Redis: &config.RedisConfig{
		MaxIdle:                3,
		IdleTimeout:            240,
		ReadTimeout:            15,
		WriteTimeout:           15,
		ConnectTimeout:         15,
		NormalTasksPollPeriod:  1000,
		DelayedTasksPollPeriod: 500,
	},
}

func InitMachinery() {
	var broker brokersiface.Broker
	var backend backendsiface.Backend
	var lock locksiface.Lock
	server := machinery.NewServer(cnf, broker, backend, lock)
	err := server.RegisterTask("sum", Sum)
	if err != nil {
		fmt.Println("注册任务失败")
	}
	worker := server.NewWorker("machinery", 1)
	go func() {
		err = worker.Launch()
		if err != nil {
			log.Println("启动worker失败")
			return
		}
	}()

	signature := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Fatal(err)
	}
	res, err := asyncResult.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("get res is %v\n", tasks.HumanReadableResults(res))

}
