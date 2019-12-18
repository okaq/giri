// okaq web serve
// human reaction time test
// AQ <aq@okaq.com>
// 2019-12-16
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "wata.html"
)

var (
	C map[string]string
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func WataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	s0 := fmt.Sprintf("%d", time.Now().UnixNano())
	w.Write([]byte(s0))
}


func main() {
	motd()
	C = make(map[string]string)
	http.HandleFunc("/", WataHandler)
	http.HandleFunc("/a", PingHandler)
	http.ListenAndServe(":8080", nil)
}


