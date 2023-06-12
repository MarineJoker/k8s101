package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	reHeader := r.Header
	rpReader := w.Header()
	rpReader = reHeader
	rpReader.Write(w)
	fmt.Fprintf(w, "这是主页")
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "访问成功")
}


func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/localhost/healthz", ok)
	http.ListenAndServe(":8080", nil)
}

