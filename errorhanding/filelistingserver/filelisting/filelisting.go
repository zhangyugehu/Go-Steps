package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const perfix = "/list/"

type userError string

func (e userError) Error() string{
	return e.Message()
}
func (e userError) Message() string{
	return string(e)
}


// HandleFileList
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, perfix) != 0{
		return userError("path must start with " + perfix)
	}
	path := request.URL.Path[len(perfix):]
	file, err:=os.Open(path)
	if err!=nil{
		return err
	}
	defer  file.Close()

	all, err:=ioutil.ReadAll(file)
	if err!=nil{
		return err
	}

	writer.Write(all)
	return nil
}

