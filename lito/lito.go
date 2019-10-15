// okaq lito
// 1024-bit array generator
// AQ <aq@okaq.com>
// 2019-10-14
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "lito.html"
)

func motd() {
	fmt.Println(time.Now().String())
}

func LitoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", LitoHandler)
	http.ListenAndServe(":8080", nil)
}


