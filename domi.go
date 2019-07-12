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


var (
	c map[string]string
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web serve localhost:8080")
}

func DomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func CacheHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// print to output stream
}

func cache() {
	c = make(map[string]string)
}

func main() {
	motd()
	// server side data cache init
	cache()
	http.HandleFunc("/", DomiHandler)
	http.ListenAndServe(":8080", nil)
}


