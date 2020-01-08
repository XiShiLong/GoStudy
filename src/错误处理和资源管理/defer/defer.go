package main

import (
	"bufio"
	"fmt"
	"os"
	"错误处理和资源管理/fib"
)

func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

type myerror struct {

}

func (myerror) Error() string {
	return "there occored an unknow error"
}

func writeFile(filename string){
	//file,err:=os.Create(filename)
	file,err:=os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	if err!=nil{
		if pathError,ok:=err.(*os.PathError);ok{
			fmt.Printf("%s, %s, %s\n",pathError.Op,
				pathError.Path,
				pathError.Err)
		}else{
			err=myerror{}
			panic(err.Error())
		}
		return
	}

	writer:=bufio.NewWriter(file)
	f:=fib.Fibonaqie()
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}

	defer file.Close()
	defer writer.Flush()
}

func main() {
	writeFile("src/错误处理和资源管理/feibo.txt")
}
