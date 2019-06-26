// ATC-SIM (WebGL / ThreeJS)
// AQ <aq@okaq.com>
// 2019-06-26
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "carto.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web serving on localhost:8080...")
}

func CartoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	// cache()
	http.HandleFunc("/", CartoHandler)
	http.ListenAndServe(":8080", nil)
}


