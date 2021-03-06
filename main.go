package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"flag"
)

/*
	simple http server which allows users to connect to pages contained inside the web directory
*/
type file struct {
	timestamp string
	content []byte
}

var caching bool
var dirToServe string
var cache map[string]file

func readFile(path string) ([]byte, error){
		if v, ok := cache[path]; ok && caching {
			// the file is cached, return the associated byte slice
			return v.content, nil
		}else{
			// open file for reading
			f, err := os.Open(path)
			if err != nil{
				return nil, err
			}
			// read the file and add it to the cache
			fmt.Println("using non chached resource", path)
			data, err := io.ReadAll(f)
			if err != nil{
				return nil, err
			}else{
				// add the file to the cache if caching is enabled
				if caching {
					cache[path] = file{"", data}
					fmt.Println("\tadded", path, "to the cache")
				}
				return  data, nil
			}
		}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// load file if exist or show error

	// if f is a directory then tries to load index.html inside the directory

	f, err := os.Open(dirToServe + string(os.PathSeparator) + r.URL.Path[1:])
	defer f.Close()
	completePath := dirToServe+string(os.PathSeparator)+r.URL.Path[1:]

	if err == nil {
		fileInfo, _ := f.Stat()
		if fileInfo.IsDir() {
			// change path requested to get index.html inside dir
			r.URL.Path = r.URL.Path + "/index.html"
			f, err = os.Open(dirToServe + string(os.PathSeparator) + r.URL.Path[1:])
			completePath += "/index.html"
		}
	}
	
	/* error 404 handling */
	if err != nil {
		content, err := readFile(dirToServe + string(os.PathSeparator) + "_errors" + string(os.PathSeparator) + "404.html")
		if err != nil {
			fmt.Fprintf(w, "ERROR 404, %s", dirToServe+string(os.PathSeparator)+"_errors"+string(os.PathSeparator)+"404.html")
		} else {
			// show the error
			fmt.Fprintf(w, "%s", content)
		}
	} else {
		
		bytes, err := readFile(completePath)

		if err != nil {
			fmt.Fprintf(w, "error while loading resource (%s)", err)
		} else {
			fmt.Fprintf(w, "%s", bytes)
		}
	}

}

func main() {
	
	var dirFlag = flag.String("d", "web", "directory to serve")
	var portFlag = flag.String("p", "80", "port on which to serve")
	var cachingFlag = flag.Bool("nc", false,  "disable caching")
	flag.Parse()
	
	dirToServe = *dirFlag
	if !(*cachingFlag) {
		fmt.Println("caching enabled")
		cache = make(map[string]file)
		caching = true
	}else{
		fmt.Println("caching disabled")
	}
	http.HandleFunc("/", handler)
	fmt.Println("currently serving", dirToServe, "on port", *portFlag)
	
	log.Fatal(http.ListenAndServe(":"+*portFlag, nil))
}
