// okaq lito
// bitmap sample and encode
// elvis gunslinger image
// for quick draw plaver v player
// human reaction time nano game
// AQ <aq@okaq.com>
// 2019-11-12
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "miko.html"
	// png bitmap data (input)
	OMG = "omg/"
	// base64 json (output)
	POW = "pow/"
)

var (
	C map[string]string
)

func motd() {
	fmt.Println("okaq web serving on localhost:8080")
	fmt.Println(time.Now().String())
}

func MikoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func cache() {
	C = make(map[string]string)
}

func main() {
	motd()
	// cache()
	http.HandleFunc("/", MikoHandler)
	http.ListenAndServe(":8080", nil)
}

// live state of file dirs
