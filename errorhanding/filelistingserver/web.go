package main

import (
	"net/http"
	"study/errorhanding/filelistingserver/filelisting"
	"os"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// wrapper func
func errWrapper(handler appHandler) func(response http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		// panic
		defer func() {
			if r:=recover(); r!=nil {
				log.Warn("Panic %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// custom error
		err:= handler(writer, request)
		if err!=nil{
			log.Warn("Error handling request: %s", err.Error())

			if userError,ok:=err.(userError);ok{
				http.Error(writer,
					userError.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsTimeout(err):
				code = http.StatusGatewayTimeout
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}














func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))


	err := http.ListenAndServe(":8083", nil)
	if err!=nil{
		panic(err)
	}
}
