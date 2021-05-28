package main

import (
	"goWebDemo/model"
	"goWebDemo/router"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
