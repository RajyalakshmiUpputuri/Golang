package handlers

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hello struct{
	l*log.Logger
}
func NewHello(l*log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Handle hello req")
}
b,err := ioutil.ReadAll(r.Body)
if err!=nil{
	h.l.Println("error reading the body",err)
	http.Error(rw,"unable to read request body",http.StatusBadRequest)
	return
}
fmt.Fprintf(rw,"Hello %s",b)
}