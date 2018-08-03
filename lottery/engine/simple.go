package engine

import (
	"study/lottery/parser"
	"study/lottery/fetcher"
	"fmt"
)

type SimpleEngine struct {
}

type Request struct {
	Url 		string
	ParserFunc 	func(byte []byte) (parser.ParseResult, error)
}


func (e *SimpleEngine) Run(request Request)  {
	byte, err := fetcher.Fetch(request.Url)
	if err != nil{
		panic(err)
	}
	result, err := request.ParserFunc(byte)
	if err != nil{
		panic(nil)
	}

	fmt.Println(result)
}