package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*func errPanic(_ http.ResponseWriter,
	_ *http.Request) error {
	panic(123)
}*/
func errprPanic(_ http.ResponseWriter,__ *http.Request)error{
	panic(123)
}

func TestErrWrapper(t *testing.T){
	tests:=[]struct{
		h appHandler
		code int
		message string
	}{
		{errprPanic,500,""},
	}

	for _,tt:=range tests{
		f:=errWrapper(tt.h)
		response:=httptest.NewRecorder()
		request:=httptest.NewRequest(
			http.MethodGet,"/",nil)
		f(response,request)
		b,_:=ioutil.ReadAll(response.Body)
		body:=string(b)
		if response.Code!=tt.code || body!=tt.message{
			t.Errorf("expect(%d , %s) got (%d, %s)",tt.code,tt.message,response.Code,body)
		}

	}
}
