// okaq nano game
// love is a stranger
// human reaction time
// AQ <aq@okaq.com>
// 2019-10-06
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "kari.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web server started on localhost:8080")
	fmt.Println("okaq lisa nano game...")
}

func KariHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	// cache()
	http.HandleFunc("/", KariHandler)
	http.ListenAndServe(":8080", nil)
}

// architect console
// pid generator
// game state cache


