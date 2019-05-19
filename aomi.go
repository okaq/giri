// webgl 2.0
// AQ <aq@okaq.com>
// 2019-05-20
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "aomi.html"
)

func motd() {
	fmt.Println("running on localhost:8080...")
	fmt.Printf("%s\n", time.Now().String())
}

func AomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", AomiHandler)
	http.ListenAndServe(":8080", nil)
}


