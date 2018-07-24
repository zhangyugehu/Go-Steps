package scheduler

import "study/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (scheduler *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (scheduler *QueuedScheduler) WorkReady(w chan engine.Request) {
	scheduler.workerChan <- w
}

func (scheduler *QueuedScheduler) Submit(r engine.Request) {
	scheduler.requestChan <- r
}

func (s *QueuedScheduler) Run()  {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r:=<-s.requestChan:
				requestQ = append(requestQ, r)
			case w:=<-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}