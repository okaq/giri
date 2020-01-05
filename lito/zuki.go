// okaq web serve
// human memory game
// AQ <aq@okaq.com>
// 2020-01-04
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "zuki.html"
	NOJI = "asobi/noto_emoji_2.json"
)

var (
	C map[string]string
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func ZukiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func NojiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,NOJI)
}

func main() {
	motd()
	// rng()
	// cache()
	http.HandleFunc("/", ZukiHandler)
	http.HandleFunc("/a", NojiHandler)
	http.ListenAndServe(":8080", nil)
}


