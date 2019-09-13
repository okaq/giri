// okaq web server
// AQ <aq@okaq.com>
// 2019-08-28
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "hiki.html"
	EMOJI = "img"
)

func motd() {
	fmt.Println(time.Now().String())
}

func cache() {
	// init memory store
	// map[key]value strings
	// generate stats for user metadata
	// visualize peer network
}

func HikiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", HikiHandler)
	fs := http.FileServer(http.Dir(EMOJI))
	http.Handle("/emoji/", http.StripPrefix("/emoji/", fs))
	// no req output, no header application/json
	// requires custom file server handler
	http.ListenAndServe(":8080", nil)
}


