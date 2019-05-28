package filelisting

import (
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)

const prefix3 = "/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList3(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix3) != 0 {
		//return errors.New("Path must start " + "with" + prefix3)
		return userError("Path must start " + "with" + prefix3)
	}
	path := request.URL.Path[len(prefix):] //  /list/google_learngo/
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		/*http.Error(writer, err.Error(), http.StatusInternalServerError)
		return*/
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
