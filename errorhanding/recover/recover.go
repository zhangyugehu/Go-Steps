package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok:=r.(error); ok{
			fmt.Println("Error occurred", err)
		}else{
			panic(r)
		}
	}()


	//a:=5
	//b:=0
	//fmt.Println(a/b)

	panic(123)
	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()
}
