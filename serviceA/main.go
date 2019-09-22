package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	// https://stackoverflow.com/a/48250354/9077800
	done := make(chan bool)
	go http.ListenAndServe(":4001", nil)
	fmt.Println("FEDERATION_SIGNAL_OK")
	fmt.Println("Server started at port 4001")
	<-done

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello service A")
}
