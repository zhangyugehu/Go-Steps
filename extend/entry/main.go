package main

import (
	"study/extend/clazz"
	"fmt"
)

func main() {
	son := clazz.Son{
		clazz.Base{"base"},
		"son",
	}

	fmt.Println(son.String())
	//son.Print()
}
