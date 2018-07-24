package engine

type ConcurrentEngine struct {
	Scheduler 	Scheduler
	WorkerCount int
	ItemChan	chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}


func (engine *ConcurrentEngine) Run(seeds ...Request)  {

	out := make(chan ParseResult)

	engine.Scheduler.Run()

	for i:=0;i<engine.WorkerCount;i++{
		createWorker(engine.Scheduler.WorkChan(), out, engine.Scheduler)
	}

	for _,r:=range seeds{
		engine.Scheduler.Submit(r)
	}

	//itemCount := 0
	for{
		result := <-out
		for _, item :=range result.Items{
			//log.Printf("Got item #%d by worker(%v): %v", itemCount, engine.Scheduler.WorkChan(), item)
			//itemCount ++

			go func() {
				engine.ItemChan <- item
			}()
		}

		for _,request:=range result.Requests{
			if isDuplicate(request.Url){
				continue
			}
			engine.Scheduler.Submit(request)
		}


	}
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrls[url]{
		return true
	}
	visitedUrls[url]=true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for{
			// tell scheduler i'm ready
			ready.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err!=nil{
				continue
			}
			out <- result
		}
	}()
}