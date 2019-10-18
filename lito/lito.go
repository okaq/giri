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

func insp() {
	// read file dir info
	// create json listings for browser
}

func LitoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// sync state to browser
}

func main() {
	motd()
	insp()
	http.HandleFunc("/", LitoHandler)
	http.HandleFunc("/a", TimeHandler)
	http.ListenAndServe(":8080", nil)
}


