package main

import(
	"net/http"
	"path/filepath"
)


func main() {
	p, _ := filepath.Abs(filepath.Dir("/home/zt/"))
	http.Handle("/", http.FileServer(http.Dir(p)))
	http.ListenAndServe(":8000", nil)
}
