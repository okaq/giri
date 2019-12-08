// okaq web serve
// bitmap sample and encode base64 json
// AQ <aq@okaq.com>
// 2019-12-06
package main

import (
	"encoding/json"
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
	// json encoded png file list
	J []byte
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
	/* debug
	for i0, f1 := range f0 {
		fmt.Printf("file: %s, item: %d\n", f1.Name(), i0)
		s0 := fmt.Sprintf("%s%s", PNG, f1.Name())
		fmt.Printf("path: %s\n", s0)
		s1 := strconv.FormatInt(f1.Size(), 10)
		s2 := fmt.Sprintf("%s:%s", s1, f1.ModTime().String())
		fmt.Printf("size:time::%s\n", s2)
	}
	*/
	for _, f1 := range f0 {
		k0 := fmt.Sprintf("%s%s", PNG, f1.Name())
		s0 := strconv.FormatInt(f1.Size(), 10)
		v0 := fmt.Sprintf("%s:%s", s0, f1.ModTime().String())
		P[k0] = v0
	}
	// fmt.Println(P)
}

func compact() {
	// encode png file list map
	// json for http writer
	var err error
	J, err = json.Marshal(P)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(J))
}

func main() {
	motd()
	// load png file list
	P = make(map[string]string)
	pictures()
	compact()
	http.HandleFunc("/", RoloHandler)
	http.HandleFunc("/a", PngHandler)
	http.ListenAndServe(":8080", nil)
}


