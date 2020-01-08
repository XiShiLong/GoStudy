package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"错误处理和资源管理/mytest-feibolistserver/filelistling"
)

type  userError interface{
	error
	Message()string
}

type usererror string

func (obj usererror) Error() string {
	return obj.Error()
}

func (obj usererror) Message() string {
	return string(obj)
}

type Handle func(writer http.ResponseWriter,request *http.Request)error

func GlobalErrDel(handle Handle)func(w http.ResponseWriter,r *http.Request){
	return func(writer http.ResponseWriter,request *http.Request){
		defer func(){
			r:=recover()
			if err,ok:=r.(userError);ok==true{
				fmt.Println(err.Error())
				http.Error(writer,err.Message(),http.StatusBadRequest)
			}

		}()
		err:=handle(writer,request)
		code:=http.StatusOK
		if err!=nil{
			log.Print("%s",err.Error())
			switch{
			case os.IsNotExist(err):
				code=http.StatusNotFound
			case os.IsPermission(err):
				code=http.StatusForbidden
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main(){
	http.HandleFunc("/",GlobalErrDel(filelistling.Handle))
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		panic(err)
	}
}
