// Hersey Aesthitics
// Logo Concepts
// AQ <aq@okaq.com>
// 2019-07-15
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "erse.html"
)

func ErseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web start on localhost:8080...")
}

func main() {
	motd()
	http.HandleFunc("/", ErseHandler)
	http.ListenAndServe(":8080", nil)
}


