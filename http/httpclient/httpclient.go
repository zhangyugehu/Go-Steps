package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	resp,err:=http.Get("http://www.baidu.com")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s", s)
}
