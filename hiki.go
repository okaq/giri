// okaq web server
// AQ <aq@okaq.com>
// 2019-08-28
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "hiki.html"
)

func motd() {
	fmt.Println(time.Now().String())
}

func HikiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", HikiHandler)
	http.ListenAndServe(":8080", nil)
}


