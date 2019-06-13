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

var (
	// global cache, accessed via channel
	C map[string]string
	W chan w
)

type w struct {
	key string
	val string
	r chan bool
}

func motd() {
	fmt.Println("running on localhost:8080...")
	fmt.Printf("%s\n", time.Now().String())
}

func cache() {
	C = make(map[string]string)
	W = make(chan w)
	go func() {
		// recievers
		w0 := <-W
		C[w0.key] = w0.val
		w0.r <- true
	}()
}

func AomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w, r, INDEX)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// parse json {key:val} data
	// send to channel for global state
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", AomiHandler)
	http.HandleFunc("/a", DataHandler)
	http.ListenAndServe(":8080", nil)
}


