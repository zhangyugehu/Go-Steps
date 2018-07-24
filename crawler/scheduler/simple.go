package scheduler

import "study/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (scheduler *SimpleScheduler) WorkChan() chan engine.Request {
	return scheduler.workerChan
}

func (scheduler *SimpleScheduler) WorkReady(chan engine.Request) {
	// do nothing
}

func (scheduler *SimpleScheduler) Run() {
	scheduler.workerChan = make(chan engine.Request)
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	//scheduler.workerChan <- request
	go func() {scheduler.workerChan <- request}()
}


