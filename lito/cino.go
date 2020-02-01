// okaq human play lab
// global psychism assessment nano game
// AQ <aq@okaq.com>
// 2020-02-01
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "cino.html"
)

var (
	C map[string]string
)

func motd() {
	fmt.Println("start okaq web serve on loaclhost:8080")
	fmt.Println(time.Now().String())
}

func CinoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	// cache
	C = make(map[string]string)
	http.HandleFunc("/", CinoHandler)
	http.ListenAndServe(":8080", nil)
}

// psychic test
// set data for peer net
// session time life
// psychic matches, leaderboard, global stats


