package main

import (
	"gpmgo/gopm/modules/log"
)

func tryRecover(){
	defer func(){
		r:=recover()
		if err,ok:=r.(error);ok==true{
			log.Error("haha "+err.Error())
		}else{
			panic("I don't know what to do")
		}
	}()
	panic(123)
}

func main(){
	tryRecover()
}
