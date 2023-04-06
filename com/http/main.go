package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *Student) toString() string {
	age := s.Age
	return "{name=" + s.Name + ",age=" + strconv.FormatInt(int64(age), 10) + "}"
}

func main() {
	sayHello()
	http.HandleFunc("/helloWorld", func(writer http.ResponseWriter, request *http.Request) {
		sayHello()
	})
	http.HandleFunc("/helloWorld2", func(writer http.ResponseWriter, request *http.Request) {
		sayHello2()
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		tar := make([]byte, request.ContentLength)
		_, err := request.Body.Read(tar)
		if err != nil {
			fmt.Println(request)
		}
		ss := bytes.NewBuffer(tar).String()
		println(ss)
		var v Student
		json.Unmarshal([]byte(ss), &v)
		v.Age++
		writer.Write(bytes.NewBufferString(v.toString()).Bytes())
	})
	http.ListenAndServe(":8080", nil)
}

func sayHello() {
	println("hello")
}

func sayHello2() {
	println("hello2")
}
