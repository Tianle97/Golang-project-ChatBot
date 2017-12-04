package main

import (
	"net/http"
	"fmt"
	"./chat"
)

// main
func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/chat/", chatHandler)
	http.ListenAndServe(":9090", nil)
} 

// chatHandler
func chatHandler(w http.ResponseWriter, r * http.Request) {
	input := r.URL.Query().Get("input")
	// takes the input we got from the request, and sends it to the Ask function
	answer := chat.Ask(input) 
	// writes the result back into the ResponseWriter
	fmt.Fprintf(w, answer)    

}
