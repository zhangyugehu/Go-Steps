package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

func errPanic(writer http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type userTestError string

func (e userTestError) Error() string{
	return e.Message()
}
func (e userTestError) Message() string{
	return string(e)
}
func errUserError(writer http.ResponseWriter, r *http.Request) error {
	return userTestError("userTestError")
}

var tests=[]struct{
	h appHandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "userTestError"},
}

func TestErrWrapper(t *testing.T) {

	for _,tt:=range tests{
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
		f(response, request)

		b, _:=ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body!=tt.message{
			t.Errorf("expect (%d, %s); got (%d, %s)",
				tt.code, tt.message, response.Code, body)
		}
	}
}

func TestErrWrapperInServer(t *testing.T)  {
	for _,tt:=range tests{
		f:=errWrapper(tt.h)
		server:=httptest.NewServer(http.HandlerFunc(f))
		resp,_:=http.Get(server.URL)
		b,_:=ioutil.ReadAll(resp.Body)
		body:=strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body!=tt.message{
			t.Errorf("expect (%d, %s); got (%d, %s)",
				tt.code, tt.message, resp.StatusCode, body)
		}
	}
}
