// okaq web serve
// human memory test
// animation timeline logic
// AQ <aq@okaq.com>
// 2020-01-15
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "baku.html"
)

func motd() {
	fmt.Println("web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func BakuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", BakuHandler)
	http.ListenAndServe(":8080", nil)
}


