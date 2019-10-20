// okaq lito
// 1024-bit array generator
// AQ <aq@okaq.com>
// 2019-10-14
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
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


