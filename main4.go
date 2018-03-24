package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"time"

	"github.com/mholt/binding"
)

type MyBinder map[string]string

func (t MyBinder) Bind(fieldsName string, values []string) error {
	t["formData"] = values[0]
	return nil
}

type FileObject struct {
	Group    string
	Data     []byte
	Content  *multipart.FileHeader
	Settings MyBinder
}

func (fo *FileObject) FieldMap(request *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&fo.Group:    "group",
		&fo.Content:  "file",
		&fo.Settings: "settings",
	}
}

func (fo *FileObject) Validate(request *http.Request) error {
	if fo.Group == "test" {
		return binding.Errors{
			binding.NewError([]string{"message"}, "forbidden", "forbidden"),
		}
	}
	return nil
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		return
	}
	obj := new(FileObject)
	obj.Settings = MyBinder{}
	if err := binding.Bind(request, obj); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("group %s\n", obj.Group)
	fmt.Printf("group %s\n", obj.Content.Filename)
	writer.Write([]byte("hello"))
}
func handler(writer http.ResponseWriter, request *http.Request) {
	multipartReader, err := request.MultipartReader()
	if err != nil {
		fmt.Printf("multipart reader errir %s", err)
	}
	obj := new(FileObject)
	for {
		reader, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		name := reader.FormName()
		if name == "group" {
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				break
			}
			obj.Group = string(data)
		} else if name == "file" {
			obj.Data, err = ioutil.ReadAll(reader)
		}
	}
	fmt.Printf("file object %s", obj)

	writer.Write([]byte("hello"))
}

func main() {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)
	mux.HandleFunc("/demo2", handler2)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  time.Second * 120,
		WriteTimeout: time.Second * 120,
	}

	err = server.Serve(listener)
	fmt.Printf("listen err %s", err)
}
