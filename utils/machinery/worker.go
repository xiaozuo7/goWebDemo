package machinery

import (
	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	"github.com/RichardKnop/machinery/v2/config"
	locksiface "github.com/RichardKnop/machinery/v2/locks/iface"
	"goWebDemo/utils"
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
}

func InitMachinery() {
	var broker brokersiface.Broker
	var backend backendsiface.Backend
	var lock locksiface.Lock
	server := machinery.NewServer(cnf, broker, backend, lock)
	server.NewWorker("machinery", 10)
}
