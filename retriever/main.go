package main

import (
	"study/retriever/mock"
	"fmt"
	"study/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever){
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Constants", v.Constants)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r=mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	//fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}

	inspect(r)
	//fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))



	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Constants)
	}else{
		fmt.Println("not a mockRetriever.")
	}
}
