// okaq web serve
// bauhaus face profile generator
// image render in browser
// andsave to disk as png file
// AQ <aq@okaq.com>
// 2019-12-15
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "trolo.html"
	PNG = "noe/"
)

var (
	// cache
	C map[string]string
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func cache() {
	C = make(map[string]string)
	// create chan workers for cache data access
}

func TroloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PngHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// save png file in request body to disk
}

func main() {
	motd()
	cache()
	// rng()
	http.HandleFunc("/", TroloHandler)
	http.HandleFunc("/a", PngHandler)
	http.ListenAndServe(":8080", nil)
}


