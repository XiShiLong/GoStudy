package mock

import (
	"fmt"
	"net/http"
	"time"
)

type Retriever struct{
	UserAgent string
	TimeOut time.Duration
	Contens string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}",r.Contens)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contens=form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string)string{
	return r.Contens
}



func (r  Retriever) Get1(url string) string {
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

