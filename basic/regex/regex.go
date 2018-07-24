package main

import (
	"regexp"
	"fmt"
)

const text = `My email is 1021091945@QQ.COM
zhangyugehu@gmail.com
18656561827
haha@163.com@abc.com
lalala  @123.com
123@hotmail.com.cn
`

func main() {

	//re, err := regexp.Compile("1021091945")
	//if err!=nil{
	//	panic(err)
	//}

	re := regexp.MustCompile(`([0-9A-Za-z]+)@([0-9A-Za-z]+)(\.[0-9A-Za-z.]+)`)

	//matchArr := re.FindAllString(text, -1)

	submatchArr := re.FindAllStringSubmatch(text, -1)

	//fmt.Println(submatchArr)

	for _,m:=range submatchArr{
		for j,n:=range m{
			fmt.Println(j, n)
		}
		//fmt.Println(m)
	}
}
