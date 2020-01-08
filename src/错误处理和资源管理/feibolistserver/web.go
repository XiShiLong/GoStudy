package main

import (
	"gpmgo/gopm/modules/log"
	"net/http"
	"os"
	"错误处理和资源管理/feibolistserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,request *http.Request) error

func errWrapper(handler appHandler)func(w http.ResponseWriter,r *http.Request){
	return func(writer http.ResponseWriter,request *http.Request){
		err:=handler(writer,request)
		if err!=nil{
			log.Warn("Error hanling request : %s",err.Error())
			code:=http.StatusOK
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
	http.HandleFunc("/list/",errWrapper(filelisting.HandleFileListing))
	err:=http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}
}
