package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

func tryDefer(){
	i := 0;
	defer fmt.Printf("i: %d \n", i)
	i++
	fmt.Printf("i: %d \n", i)
	i++
	fmt.Printf("i: %d \n", i)
	i++
	fmt.Printf("i: %d \n", i)
}

func writeFile(filename string) {
	//file, err:=os.Create(filename)
	file, err:=os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("custom error")
	if err != nil{
		//fmt.Println(err.Error())
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		}else{
			fmt.Printf("%s, %s, %s.", pathError.Err,pathError.Op,pathError.Path)
		}
	}
	defer file.Close()

	writer:=bufio.NewWriter(file)
	defer writer.Flush()

	for i:=0;i<100;i++ {
		fmt.Fprintln(writer, "writer ", i)
	}

}

func main() {

	//tryDefer()

	writeFile("defer.txt")
}
