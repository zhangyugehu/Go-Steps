package main

import "fmt"

func defineArr(){
	var arr1 [5]int;
	arr2 := [3]int{1,2,3}
	arr3:= [...]int {1,2,3,4,5}
	var grid [4][5]int // 四行五列

	fmt.Println(arr1, arr2 , arr3)
	fmt.Println(grid)
}

func traverseArr(){
	arr:= [...]string {"a","b","c","d","e"}

	//func 1:
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	// func 2:
	for i,v:=range arr{
		fmt.Println(i, v)
	}
}

func printArr(arr [3]int){
	arr[0] = 100
	for i:=range arr{
		fmt.Println(arr[i])
	}
}

func main() {
	//defineArr()

	//traverseArr()

	// 数组参数类型是值传递
	//arr := [3]int{1,2,3}
	//printArr(arr)
	//fmt.Println(arr)


}
