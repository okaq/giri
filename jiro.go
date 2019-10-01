// dont you forget about me
// play proto
// AQ <aq@okaq.com>
// 2019-09-18
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "jiro.html"
)

func motd() {
	fmt.Println("dont you forget about me")
	fmt.Println("starting web server on localhost:8080")
	fmt.Println(time.Now().String())
}

func cache() {
	// format of cache
	// map[uint64]interface
	// strings as sequential naturals
	// values can be any thing, any size
}

func JiroHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// pid format
	// clientRand:clientTimestamp::serverRand:serverTimestamp
}

func CacheHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// all cache hits are updates
	// no history, only live global state
}

func main() {
	motd()
	cache()
	http.HandleFunc("/", JiroHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", CacheHandler)
	http.ListenAndServe(":8080", nil)
}


