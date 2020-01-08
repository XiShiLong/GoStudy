package main

import (
	"fmt"
	"time"
	mock "接口编程/realll"
)

const url="http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func download(r Retriever)string{
	return r.Get(url)
}

func post(poster Poster){
	poster.Post(url,
		map[string]string{
			"name":"ccmouse",
			"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster)string{

	s.Post(url,map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main(){
	r:=mock.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
		Contens:"this .is a fake imooc.com",
	}

	fmt.Println("try a session")
	s:=session(&r)
	fmt.Println(s)
	fmt.Println(r.String())
}
