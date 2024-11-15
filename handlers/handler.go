package handlers

import (
	"net/http"
	"regexp"
	"strconv"
)

var API = "https://groupietrackers.herokuapp.com/api"

type ArtistHanlder struct{}

var (
	Artist        []artist
	artistUrls, _ = regexp.Compile(`^artists\?id=\d+$`)
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	FetchData(&Artist)
	renderTemplate(w, "index.html", Artist)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		dates    Date
		location Location
		relation Relation
		artist   artist
	)

	user_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if Artist != nil {
		_, ok := Artist[user_id-1].Locations.(string)
		if ok {
			FetchData(&Artist[user_id-1].Locations, Artist[user_id-1].Locations.(string))
			FetchData(&Artist[user_id-1].ConcertDates, Artist[user_id-1].ConcertDates.(string))
			FetchData(&Artist[user_id-1].Relations, Artist[user_id-1].Relations.(string))
		}
		renderTemplate(w, "artist.html", Artist)

	} else {
		api := API + "/artists/" + strconv.Itoa(user_id)
		FetchData(&artist, api)
		_, ok := artist.Locations.(string)
		if ok {
			FetchData(&location, artist.Locations.(string))
			FetchData(&relation, artist.Relations.(string))
			FetchData(&dates, artist.ConcertDates.(string))
			artist.ConcertDates = dates
			artist.Locations = location
			artist.Relations = relation
		}
		renderTemplate(w, "artist.html", artist)
	}
}

// There is a problem in this demarch
// while we can use the function to fetch the data
// But the the first way is more useful

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, "error.html", http.StatusNotFound)
}

func (A ArtistHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && artistUrls.MatchString(r.URL.Path[1:]+"?"+r.URL.RawQuery):
		GetUser(w, r)
	case r.Method == http.MethodGet && (r.URL.Path == "/" || r.URL.Path[1:] == "artists"):
		GetUsers(w, r)
	default:
		NotFound(w, r)
	}
}
