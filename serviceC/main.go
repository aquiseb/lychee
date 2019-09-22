package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	// https://stackoverflow.com/a/48250354/9077800
	done := make(chan bool)
	go http.ListenAndServe(":4000", nil)
	fmt.Println("ok")
	fmt.Println("Server started at port 4000")
	<-done
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello service C")
}
