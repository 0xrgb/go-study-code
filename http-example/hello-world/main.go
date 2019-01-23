package main

import (
	"fmt"
	"log"
	"net/http"
)

// `/`로 시작하는 모든 주소에 대해서 처리하는 함수.
func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL: %q\n", req.URL)
}

// `/hello`로 시작하는 특정 주소에 대해서 처리하는 함수.
func helloHandler(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/hello/"):]
	if len(name) == 0 {
		name = "world"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello/", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
