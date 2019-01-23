// static 폴더 안에 파일을 넣고 실행한다.
package main

import (
	"log"
	"net/http"
)

// 로그를 작성하는 Middleware
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s => %q", req.RemoteAddr, req.URL)
		h.ServeHTTP(w, req)
	})
}

func main() {
	h := http.NewServeMux()
	h.Handle("/", http.FileServer(http.Dir("static")))

	hl := logger(h)
	err := http.ListenAndServe(":8080", hl)
	if err != nil {
		log.Fatal(err)
	}
}
