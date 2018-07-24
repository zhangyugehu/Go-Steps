package clazz

import "fmt"

type Base struct {
	Name string
}

func (this *Base) Print(){
	fmt.Println("Base ", this.Name)
}

type Son struct {
	Base
	SonName string
}

func (this *Son) String() string {
	return "Name:"+this.Name +
		", SonName:" + this.SonName;
}

func (this *Son) Print(){
	fmt.Println("Son", this.Name, this.SonName)
}

