// okaq lito
// 1024-bit array generator
// AQ <aq@okaq.com>
// 2019-10-14
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	INDEX = "lito.html"
	OMG = "omg/"
	POW = "pow/"
)

var (
	List []byte
	Cache map[string]string
)

func motd() {
	fmt.Println(time.Now().String())
}

func insp() {
	// read file dir info
	f, err := ioutil.ReadDir(OMG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(f))
	// create json listings for browser

	// pretty print file info
	// marshal to json for browser

	for i, f0 := range f {
		fmt.Println(i)
		fmt.Println(f0.Name())
		fmt.Println(f0.Size())
		fmt.Println(f0.Mode())
		fmt.Println(f0.ModTime())
		fmt.Println(f0.IsDir())
		fmt.Println()
		// json inline struct
	}

	// alt: just marshal []FileInfo
	b, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	// achtung! requires pointer deref

	f1 := make([]os.FileInfo, len(f))
	for i, f0 := range f {
		f1[i] = f0
	}
	fmt.Println(f1)
	b1, err := json.Marshal(f1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b1)
	fmt.Println(len(b1))
	fmt.Println(string(b1))
	// empty!

	// just try a map
	f2 := make(map[string]string)
	for i, f0 := range f {
		// s0 := string(i)
		s0 := strconv.Itoa(i)
		i1 := strconv.FormatInt(f0.Size(), 10)
		s1 := fmt.Sprintf("%s|%s|%s", f0.Name(), i1, f0.ModTime().String())
		f2[s0] = s1
	}
	fmt.Println(f2)
	b2, err := json.Marshal(f2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b2)
	fmt.Println(len(b2))
	fmt.Println(string(b2))
	// set List for output writing
	List = b2
}

func data() {
	// init cache
	Cache = make(map[string]string)
	// pop pid zero
	// sync chan
}

func LitoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// sync state to browser
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Content-type", "application/json")
	w.Write(List)
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// output cache to http writer
	// stat cache includes
	// req count, player total, pid live
}

func main() {
	motd()
	insp()
	data()
	http.HandleFunc("/", LitoHandler)
	http.HandleFunc("/a", TimeHandler)
	http.HandleFunc("/b", ListHandler)
	http.HandleFunc("/c", StatHandler)
	http.ListenAndServe(":8080", nil)
}


