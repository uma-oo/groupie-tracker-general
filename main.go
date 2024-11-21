package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"groupie/handlers"
)

var Templates *template.Template

func init(){
  var err error
   Templates, err = template.ParseGlob("./templates/*.html")
   if err!= nil {
      log.Fatal("The files that need to be")
   }
}




func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handlers.ArtistHanlder{})
	fmt.Println("Listening on http://localhost:3535")
	http.ListenAndServe(":3535", mux)
}
