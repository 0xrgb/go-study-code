package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var counter int64

// `/counter`로 시작하는 모든 주소에 대해서 처리하는 함수.
func counterHandler(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("visited")

	// 쿠키 여부에 따라 처리
	if err != nil {
		// 쿠키가 없으므로 counter에 1을 더하고 쿠키를 만들어준다.
		cookie := http.Cookie{
			Name:    "visited",
			Value:   "true",
			Expires: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, &cookie)

		now := atomic.AddInt64(&counter, 1)
		fmt.Fprintf(w, "Counter: %d\nCookie: No\n", now)
	} else {
		now := atomic.LoadInt64(&counter)
		fmt.Fprintf(w, "Counter: %d\nCookie: Yes\n", now)
	}
}

func main() {
	http.HandleFunc("/counter/", counterHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
