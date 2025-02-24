package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	//m := sync.Mutex{}
	//m.Lock()
	atomic.AddUint64(&number, 1)
	//m.Unlock()
	time.Sleep(600 * time.Microsecond)
	w.Write([]byte(fmt.Sprintf("You are the visitant of number %d", number)))

}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":9000", nil)

}
