//Tianle Shu
//G00353418

package main

import (

	"fmt"
	"net/http"
	"./chat"
    "log"
)

func HandleResponse(w http.ResponseWriter, request *http.Request) {
	text := request.URL.Query().Get("text")
	answer := chat.RespondTo(text) // takes the text we got from the request, and sends it to the Ask function
	fmt.Fprintf(w, answer)    // writes the result back into the ResponseWriter

}

func main() {

			dir := http.Dir("./image")  //image folder

			fileServer := http.FileServer(dir)

			http.HandleFunc("/", fileServer)

			http.HandleFunc("/conversation", HandleResponse)

			// Start the server at http://localhost:9090
			log.Fatal(http.ListenAndServe(":9090", nil))
		}