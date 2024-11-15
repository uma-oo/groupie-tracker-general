package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmp string, data interface{}) {
	temp, err := template.ParseFiles("./templates/" + tmp)
	if err != nil {
		fmt.Println()
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "<h1>Internal Server Error 500</h1>")
		return
	}
	err = temp.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "<h1>Internal Server Error 500</h1>")
		return
	}
}

// working with maps here but if it has to be a struct it will be more efficient
// for later

func GetApis() map[string]string {
	var map_api map[string]string
	resp, err := http.Get(API)
	if err != nil {
		log.Fatal("There was an error while fetching the main API!")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error while Reading the body of the response!!")
	}

	err = json.Unmarshal(data, &map_api)
	if err != nil {
		log.Fatal("ERROR HERE IN GETAPIS", err)
	}
	return map_api
}

func FetchUrl[T any](holder *T, url string) {
}

func FetchData[T any](holder *T, name_api ...string) { // important
	// artists ghadi nkhdmuh by default ila la
	// we will fetch the url 3adii
	url_fetched := "artists"
	api := GetApis()[url_fetched]
	if len(name_api) == 1 {
		url_fetched = name_api[0]
		api = url_fetched
	}

	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(holder)
	if err != nil {
		panic(err)
	}
}
