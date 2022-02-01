package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
	simple http server which allows users to connect to pages contained inside the web directory
*/

var dirToServe string

func handler(w http.ResponseWriter, r *http.Request) {
	// load file if exist or show error

	// if f is a directory then tries to load index.html inside the directory

	f, err := os.Open(dirToServe + string(os.PathSeparator) + r.URL.Path[1:])
	defer f.Close()

	if err == nil {
		fileInfo, _ := f.Stat()
		if fileInfo.IsDir() {
			// change path requested to get index.html inside dir
			r.URL.Path = r.URL.Path + "/index.html"
			f, err = os.Open(dirToServe + string(os.PathSeparator) + r.URL.Path[1:])
		}
	}

	/* error 404 handling */
	if err != nil {
		f, err := os.Open(dirToServe + string(os.PathSeparator) + "_errors" + string(os.PathSeparator) + "404.html")
		if err != nil {
			fmt.Fprintf(w, "ERROR 404, %s", dirToServe+string(os.PathSeparator)+"_errors"+string(os.PathSeparator)+"404.html")
		} else {
			bytes, err := io.ReadAll(f)
			if err == nil {
				// show the error
				fmt.Fprintf(w, "%s", bytes)
			}
		}
	} else {
		bytes, err := io.ReadAll(f)

		if err != nil {

			fmt.Fprintf(w, "error while loading resource (%s)", err)
		} else {
			fmt.Fprintf(w, "%s", bytes)
		}
	}

}

func main() {
	if len(os.Args) > 2 {
		dirToServe = os.Args[1]
	} else {
		dirToServe = "web"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
