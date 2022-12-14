package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func appHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Now(), "Hello from my new fresh server")
	w.Write([]byte("<h1>Hello from my new fresh server!</h1>"))

}

func main() {
	http.HandleFunc("/", appHandler)

	log.Println("Started, serving on port 80")
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
