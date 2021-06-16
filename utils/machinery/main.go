package main

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/example/tracers"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"github.com/urfave/cli"
	"goWebDemo/utils"
	"goWebDemo/utils/machinery/exampletasks"
	"log"
	"os"
	"time"
)

var app *cli.App

// 初始化app版本
func init() {
	app = cli.NewApp()
	app.Name = "machinery"
	app.Usage = "machinery异步任务"
	app.Version = "0.0.1"
}

// 主函数入口
func main() {
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "加载machinery worker",
			Action: func(c *cli.Context) error {
				if err := worker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},

		{
			Name:  "send",
			Usage: "发送task",
			Action: func(c *cli.Context) error {
				if err := send(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}
	_ = app.Run(os.Args)
}

// 启动服务
func startServer() (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:          utils.Broker,
		DefaultQueue:    "machinery_task",
		ResultBackend:   "redis://" + utils.RedisBackendAuth + "@" + utils.RedisBackendHost + ":" + utils.RedisBackendPort,
		ResultsExpireIn: 3600,
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "machinery_task",
			PrefetchCount: 3,
		},
	}

	broker := amqpbroker.New(cnf)
	backend := redisbackend.NewGR(cnf, []string{utils.RedisBackendAuth + "@" + utils.RedisBackendHost + ":" + utils.RedisBackendPort}, utils.RedisBackendDB)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	registerTasks := map[string]interface{}{
		"add": exampletasks.Add,
	}
	return server, server.RegisterTasks(registerTasks)

}

// 初始化worker
func worker() error {
	consumerTag := "machinery_worker"
	cleanUp, err := tracers.SetupTracer(consumerTag)
	if err != nil {
		log.Fatalln("无法实例化一个tracer: ", err)
	}
	defer cleanUp()

	server, err := startServer()
	if err != nil {
		return err
	}

	worker := server.NewWorker(consumerTag, 4)
	errorHandler := func(err error) {
		log.Println("错误处理:", err)
	}
	preTaskHandler := func(signature *tasks.Signature) {
		log.Println("start a task::", signature.Name)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		log.Println("task succeed", signature.Name)
	}
	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)
	return worker.Launch()

}

func send() error {
	cleanUp, err := tracers.SetupTracer("sender")
	if err != nil {
		log.Fatalln("无法实例化一个tracer: ", err)
	}
	defer cleanUp()

	server, err := startServer()
	if err != nil {
		return err
	}

	eta1 := time.Now().UTC().Add(time.Second * 5)  // 延迟5秒
	var (
		addTask0, addTask1 tasks.Signature
	)

	var initTasks = func() {
		addTask0 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{Type: "int64",
					Value: 1},
			},
			ETA: &eta1,
		}
		addTask1 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 2,
				},
				{
					Type:  "int64",
					Value: 2,
				},
			},
		}
	}
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracinglog.String("batch.id", batchID))
	log.Println("Starting batch: ", batchID)

	initTasks()

	log.Println("Single task:")
	asyncResult, err := server.SendTaskWithContext(ctx, &addTask0)
	if err != nil {
		return fmt.Errorf("不能发送task: %s", err.Error())
	}
	results, err := asyncResult.Get(time.Millisecond * 5)
	if err != nil {
		return fmt.Errorf("获取task结果失败: ", err.Error())
	}
	log.Printf("1 + 1 = %v\n", tasks.HumanReadableResults(results))

	//initTasks()
	//log.Println("Group of tasks")
	//
	//group, err := tasks.NewGroup(&addTask0, &addTask1)
	//if err != nil {
	//	return fmt.Errorf("创建群组失败: %s", err.Error())
	//}
	//asyncResults, err := server.SendGroupWithContext(ctx, group, 10)
	//if err != nil {
	//	return fmt.Errorf("不能发送群组任务: %s", err.Error())
	//}
	//
	//for _, asyncRes := range asyncResults {
	//	results, err := asyncRes.Get(time.Millisecond * 5)
	//	if err != nil {
	//		return fmt.Errorf("获取task结果失败: ", err.Error())
	//	}
	//	log.Printf("%v + %v = %v\n", asyncRes.Signature.Args[0].Value, asyncRes.Signature.Args[1].Value, tasks.HumanReadableResults(results))
	//}
	return nil
}
