package main

import (
	"fmt"
	"time"
	mock "接口编程/realll"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever)string{
	return r.Get("http://www.imooc.com")
}

func main(){
	var r Retriever
	r=&mock.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	fmt.Printf("%T || %v\n",r,r)
}
