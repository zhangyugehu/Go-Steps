package main

import (
	"study/queue"
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func testSparse(){
	s:=intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(100000)


	fmt.Println(s.Has(1000))
}

func main() {
	q:=queue.Queue{0}
	fmt.Println(q)


	testSparse()

}
