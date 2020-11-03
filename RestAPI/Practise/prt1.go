package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", prt1Handler)
	http.ListenAndServe(":8080", nil)

}
func prt1Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
