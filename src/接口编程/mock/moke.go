package mock

import (
	"fmt"
	"net/http"
)

type Retriever struct{
	Contents string
}

func (r Retriever) Get(url string) string {
	rep,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
		return "error"
	}

	data:=make([]byte,1024)

	len,_:=rep.Body.Read(data)

	if rep.Body!=nil{
		defer rep.Body.Close()
	}

	return string(data[0:len])

}

