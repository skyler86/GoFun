package main

import (
	"TodoList/config"
	"TodoList/route"
)

func main() {
	config.Init()
	r := route.NewRouter()
	_=r.Run(config.HttpPort)
}