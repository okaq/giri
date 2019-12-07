// okaq web serve
// bitmap sample and encode base64 json
// AQ <aq@okaq.com>
// 2019-12-06
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	INDEX = "rolo.html"
	// png dir
	PNG = "omg/"
	// json dir
	JSON = "pow/"
)

var (
	// png file list
	P map[string]string
)

func motd() {
	fmt.Println("okaq web serve at localhost:8080")
	fmt.Println(time.Now().String())
}

func RoloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PngHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// static png server at PNG
}

func pictures() {
	// load png files into map
	f0, err := ioutil.ReadDir(PNG)
	if err != nil {
		fmt.Println(err)
	}
	for i0, f1 := range f0 {
		fmt.Printf("file: %s, item: %d\n", f1.Name(), i0)
		s0 := fmt.Sprintf("%s%s", PNG, f1.Name())
		fmt.Printf("path: %s\n", s0)
		s1 := strconv.FormatInt(f1.Size(), 10)
		s2 := fmt.Sprintf("%s:%s", s1, f1.ModTime().String())
		fmt.Printf("size:time::%s\n", s2)
	}
}

func main() {
	motd()
	// load png file list
	pictures()
	http.HandleFunc("/", RoloHandler)
	http.HandleFunc("/a", PngHandler)
	http.ListenAndServe(":8080", nil)
}


