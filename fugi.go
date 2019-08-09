// Don't You Forget About Me
// peer net visualization
// 2019-08-09
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "fugi.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web server started localhost:8080")
}

func FugiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", FugiHandler)
	http.ListenAndServe(":8080", nil)
}


