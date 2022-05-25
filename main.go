package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)


func main() {
	fString := os.Args[1]
	length := len(fString)
	if fString[length-1] != '/' {
		fString = fString + "/"
	}
	p, _ := filepath.Abs(filepath.Dir(fString))
	fmt.Println(p)
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
