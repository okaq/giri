// dont you forget about me
// play proto
// AQ <aq@okaq.com>
// 2019-09-18
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "jiro.html"
)

func motd() {
	fmt.Println("dont you forget about me")
	fmt.Println("starting web server on localhost:8080")
	fmt.Println(time.Now().String())
}

func JiroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	// cache()
	http.HandleFunc("/", JiroHandler)
	http.ListenAndServe(":8080", nil)
}


