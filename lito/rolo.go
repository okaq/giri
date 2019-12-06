// okaq web serve
// bitmap sample and encode base64 json
// AQ <aq@okaq.com>
// 2019-12-06
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "rolo.html"
	// png dir
	PNG = "omg/"
	// json dir
	JSON = "pow/"
)

func motd() {
	fmt.Println("okaq web serve at localhost:8080")
	fmt.Println(time.Now().String())
}

func RoloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PngHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// static png server at PNG
}

func main() {
	motd()
	http.HandleFunc("/", RoloHandler)
	http.HandleFunc("/a", PngHandler)
	http.ListenAndServe(":8080", nil)
}


