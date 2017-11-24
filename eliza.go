package main

import (

	"log"
	"net/http"
	"./chat"

)

func main() {

			dir := http.Dir("./image")  //image folder

			fileServer := http.FileServer(dir)

			http.HandleFunc("/", fileServer)

			http.HandleFunc("/conversation", func(w http.ResponseWriter, r *http.Request) {
					input := r.URL.Query()Get("input")

					relacement := chat.RespondTo(input)

					fmt.Fprintf(w,replacement)
			})

			// Start the server at http://localhost:9090
			log.Fatal(http.ListenAndServe(":9090", nil))
		}