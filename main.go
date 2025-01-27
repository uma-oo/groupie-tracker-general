package main

import (
	"fmt"
	"net/http"

	"groupie/handlers"
)

func main() {
	http.HandleFunc("/assets/", handlers.HandleAssets)
	http.Handle("/", &handlers.ArtistHanlder{})
	fmt.Println("Listening on http://localhost:3535")
	http.ListenAndServe(":3535", nil)
}
