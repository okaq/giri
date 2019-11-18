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
	// Png file meta data map
	// Json byte output from marshal
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
}

func cache() {
	C = make(map[string]string)
}

func load() {
	// png file meta data
	f, err := ioutil.ReadDir(OMG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(f))
}

func map() {
	// png file key value
	f0 := make(map[string]string)
	for i, f1 := range f {
		s0 := strconv.Itoa(i)
		i0 := strconv.FormatInt(f1.Size(), 10)
		s1 := fmt.Sprintf("%s|%s|%s",f1.Name(),i1,f0.ModTime().String())
		f0[s0] = s1
	}
}

func main() {
	motd()
	// cache()
	// load()
	// map()
	http.HandleFunc("/", MikoHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/b", OmgHandler)
	http.ListenAndServe(":8080", nil)
}

// live state of file dirs
