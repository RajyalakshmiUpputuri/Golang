
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("*******Connecting*****")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Reading the body error", err)

			http.Error(rw, "Error has been found", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "About to communicate %s", d)
	})
	http.HandleFunc("Connect", func(http.ResponseWriter, *http.Request) {
		log.Println("Connected Successfully")
	})

	

	log.Println("Server is about to start")
	 http.ListenAndServe(":9090", nil)
	
}
