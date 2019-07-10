// Bacterial Tactics
// playable seeding
// 2019-07-10
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "domi.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web serve localhost:8080")
}

func DomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	// server side data cache init
	http.HandleFunc("/", DomiHandler)
	http.ListenAndServe(":8080", nil)
}


