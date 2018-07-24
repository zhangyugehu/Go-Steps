package main

import (
	"study/crawler/engine"
	"study/crawler/scheduler"
	"study/crawler/zhenai/parser"
	"study/crawler/persist"
)


func main() {
	itemSaver, err := persist.ItemSaver()
	if err!=nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: 	&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:	itemSaver,
	}

	e.Run(engine.Request{
		Url:		"http://www.zhenai.com/zhenghun",
		ParserFunc:	parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:		"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc:	parser.ParseCity,
	//})

}
