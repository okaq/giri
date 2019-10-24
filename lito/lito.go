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
	"time"
)

const (
	INDEX = "lito.html"
	OMG = "omg/"
	POW = "pow/"
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
}

func LitoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// sync state to browser
}

func main() {
	motd()
	insp()
	http.HandleFunc("/", LitoHandler)
	http.HandleFunc("/a", TimeHandler)
	http.ListenAndServe(":8080", nil)
}


