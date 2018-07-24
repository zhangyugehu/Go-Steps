package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan int  {
	c:=make(chan int)
	//go func() {
	//	for{
	//		fmt.Printf("worker %d received %c\n", id, <-c)
	//	}
	//}()
	go worker(id, c)
	return c
}

// 只能发送数据
func createSenderWorker(id int) chan<- int  {
	c:=make(chan int)
	go func() {
		for{
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}


// 只能接受数据
func createReceiverWorker(id int) <-chan int  {
	c:=make(chan int)
	go func() {
		for{
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {

	var channels [10]chan int
	var senderChannels [10]chan<- int
	var receiverChannels [10]<-chan int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		senderChannels[i] = createSenderWorker(i)
		receiverChannels[i] = createReceiverWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	
	time.Sleep(time.Microsecond)

}

func worker(id int, c chan int)   {
	for{
		//if n, ok := <-c; !ok{
		//	break
		//}else {
		//	fmt.Printf("worker %d received %d\n", id, n)
		//}

		for n:=range c{
			fmt.Printf("worker %d received %d\n", id, n)
		}
	}
}


func bufferChannel() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	c <- 'f'
	c <- 'g'
	time.Sleep(time.Microsecond)
}

func channelClose()  {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	c <- 'f'
	c <- 'g'
	close(c)
	time.Sleep(time.Microsecond)

}

func main() {
	chanDemo()

	//bufferChannel()
	//channelClose()

}
