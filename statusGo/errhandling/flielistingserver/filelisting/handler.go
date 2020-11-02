package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFilelist(writer http.ResponseWriter, request *http.Request) error {
	path:= request.URL.Path[len("/list/"):]
	file,err := os.Open(path)
	if err !=nil{
		return err
	}
	defer file.Close()
	all,err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}
	_, _ = writer.Write(all)
	return nil
}