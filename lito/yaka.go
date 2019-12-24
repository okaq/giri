// okaq web serve
// human reaction time game
// AQ <aq@okaq.com>
// 2019-12-24
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "yaka.html"
)

var (
	C map[string]string
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func rng() {
	// time seed now
}

func cache() {
	C = make(map[string]string)
}

func YakaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	rng()
	cache()
	http.HandleFunc("/", YakaHandler)
	http.ListenAndServe(":8080", nil)
}


