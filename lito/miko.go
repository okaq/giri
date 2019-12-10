// okaq lito
// bitmap sample and encode
// elvis gunslinger image
// for quick draw plaver v player
// human reaction time nano game
// AQ <aq@okaq.com>
// 2019-11-12
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"strconv"
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
	// Png file meta data map
	P map[string]string
	// Json byte output from marshal
	J []byte
	// file list
	F []os.FileInfo
)

func motd() {
	fmt.Println("okaq web serving on localhost:8080")
	fmt.Println(time.Now().String())
}

func MikoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// generate and cache pid
}

func OmgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// Png file meta data json
	w.Header().Set("Content-type","application/json")
	w.Write(J)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// usage stats output to browser

	// avg human response time per player
	// global human response time for all players
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// pretty print entire server state

	// endpoint for player response time data
	// each request represents one iteration
	// button pop, timer fires, click split time
}

func PeerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// connect the peer network
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// static file assets list
	// w.Header().Set("Content-Type","application/json")
	// w.Write(J)
}

func cache() {
	C = make(map[string]string)
}

func load() {
	// png file meta data
	var err error
	F, err := ioutil.ReadDir(OMG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(F))
}

func mapi() {
	// png file key value
	// f0 := make(map[string]string)
	P = make(map[string]string)
	for i, f1 := range F {
		s0 := strconv.Itoa(i)
		i0 := strconv.FormatInt(f1.Size(), 10)
		s1 := fmt.Sprintf("%s|%s|%s",f1.Name(),i0,f1.ModTime().String())
		// f0[s0] = s1
		P[s0] = s1
	}
}

func encode() {
	// json encode map data
	// b0, err := json.Marshal(f0)
	var err error
	J, err = json.Marshal(P)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(b0)
	fmt.Printf("length in bytes of json png file data: %d\n", len(J))
}

func main() {
	motd()
	// cache()
	// load()
	// map()
	// encode()
	http.HandleFunc("/", MikoHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", OmgHandler)
	http.HandleFunc("/c", StatHandler)
	http.HandleFunc("/d", DataHandler)
	http.HandleFunc("/e", PeerHandler)
	http.HandleFunc("/f", FileHandler)
	http.ListenAndServe(":8080", nil)
}

// live state of file dirs
