package main

import (
	"fmt"
	"sync"
)

func  doworker(id int,c chan int,wg *sync.WaitGroup){
	for n:=range c{
		fmt.Printf("worker %d receive %c \n",id,n)
		wg.Done()
	}
}

type worker struct{
	in chan int
	wg *sync.WaitGroup
}

func channelDemo(){
	var wg sync.WaitGroup

	var workers [10]worker
	for i:=0;i<10;i++{
		workers[i]=createWorker(i,&wg)
	}

	wg.Add(20)
	for i:=0;i<10;i++{
		workers[i].in<-'a'+i
	}

	for i:=0;i<10;i++{
		workers[i].in<-'A'+i
	}

	wg.Wait()
}

func createWorker(id int,wg *sync.WaitGroup)worker{
	w:=worker{
		in : make(chan int),
		wg :wg,
	}
	go doworker(id,w.in,w.wg)
	return w
}


func main(){
	channelDemo()

}