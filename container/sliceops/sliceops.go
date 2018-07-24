package main

import "fmt"

func printSlice(arr []int){
	fmt.Printf("len=%d, cap=%d\n", len(arr), cap(arr))
}

func main() {

	var s []int //Zero value for slice is nil

	for i:=0;i<10;i++{
		printSlice(s)
		s=append(s, 2*i+1)
	}
	fmt.Println(s)

	s1:=[]int{1,2,3,4,5}
	printSlice(s1)

	s2:=make([]int, 20)
	printSlice(s2)

	s3:=make([]int, 2, 100)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1);
	printSlice(s2)

	fmt.Println("Deleting elements")
	s2 = append(s2[:3], s2[:4]...)
	fmt.Println(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2=s2[1:]

	fmt.Println("Popping from back")
	tail:=s2[len(s2)-1]
	fmt.Println(tail, s2)
	s2=s2[:len(s2)-1]
	fmt.Println(front, s2)
}
