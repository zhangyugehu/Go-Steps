package maps

import "fmt"

func defineMap(){
	fmt.Println("Define Map")
	m:=map[string]string{
		"name":"zhangsan",
		"age":"12",
	}

	m2:=make(map[string]string)

	var m3 map[string]string

	fmt.Println(m,m2,m3)
}

func traverseMap(){

	fmt.Println("Travering map")
	m:=map[string]string{
		"name":"zhangsan",
		"age":"12",
	}

	for k,v:=range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting Values")
	fmt.Println("name " + m["name"])
	fmt.Println("other " + m["other"])
	if age, exist := m["age"]; exist{
		fmt.Println(age)
	}else{
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	fmt.Println(m)
}

func main() {
	defineMap()
	traverseMap();
}
