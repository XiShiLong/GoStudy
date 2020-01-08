package filelistling

import (
	"io/ioutil"
	"kataras/iris/core/errors"
	"net/http"
	"os"
	"strings"
)



func Handle(writer http.ResponseWriter,request *http.Request) error{

	if !strings.Contains(request.URL.Path,"/list/"){
		return errors.New("path do not contant '/list/'")
	}

	path:=request.URL.Path[len("/list/"):]

	file,err:=os.Open(path)
	if err!=nil{
		return err
	}
	content,err:=ioutil.ReadAll(file)
	if err!=nil{
		return err
	}

	_,err2:=writer.Write(content)
	if err2!=nil{
		return err
	}
	return nil
}
