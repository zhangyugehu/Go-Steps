package main

import "fmt"

func slice(){
	arr:=[...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[:] = ", arr[:])
}

func updateSlice(arr []int){
	arr[0] = 100
}

func main() {
	//slice()
	arr:=[...]int{0,1,2,3,4,5,6,7}
	s1:=arr[2:]
	s2:=arr[:]
	fmt.Println("s1", s1)
	updateSlice(s1)
	fmt.Println("After updateSlice(s1)", s1)

	fmt.Println("s2", s2)
	updateSlice(s2)
	fmt.Println("After updateSlice(s2)", s2)
	s2=s2[2:]
	fmt.Println("ReSlice", s2)

	//slice可以向后扩展，不可以向前扩展
	s3:=arr[2:6]
	s4:=s3[3:5]
	fmt.Println(arr, s3, s4)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))

	// append
	s5:=append(s3,10)
	s6:=append(s5, 11)
	// 继续添加后数组长度超过arr，系统将重新分配数组
	s7:=append(s6, 12)
	fmt.Println(s3, s5, s6, s7, arr)
}
