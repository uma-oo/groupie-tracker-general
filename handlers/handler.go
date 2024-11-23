package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

var API = "https://groupietrackers.herokuapp.com/api"

type ArtistHanlder struct{}

var (
	Artist        []artist
	artistUrls, _ = regexp.Compile(`^artists\?id=\d*`)
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var wait_group sync.WaitGroup
	wait_group.Add(1)
	go FetchData(&Artist, &wait_group)
	wait_group.Wait()
	renderTemplate(w, "Artists.html", Artist, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		dates      Date
		location   Location
		relation   Relation
		artist     artist
		wait_group sync.WaitGroup
	)
	user_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if r.URL.Query().Get("id") == "" {
		fmt.Println("here")
		renderTemplate(w, "error.html", http.StatusBadRequest, http.StatusBadRequest)
		return
	}
	if Artist != nil {
		if user_id > len(Artist) || user_id <= 0 {
			renderTemplate(w, "error.html", http.StatusNotFound, http.StatusNotFound)
			return
		}
		_, ok := Artist[user_id-1].Locations.(string)
		if ok {
			wait_group.Add(3)
			go FetchData(&location, &wait_group, Artist[user_id-1].Locations.(string))
			go FetchData(&dates, &wait_group, Artist[user_id-1].ConcertDates.(string))
			go FetchData(&relation, &wait_group, Artist[user_id-1].Relations.(string))
			wait_group.Wait()
			Artist[user_id-1].ConcertDates = dates
			Artist[user_id-1].Locations = location
			Artist[user_id-1].Relations = relation
		}
		renderTemplate(w, "artist.html", Artist[user_id-1], http.StatusOK)

	} else {
		api := API + "/artists/" + strconv.Itoa(user_id)
		wait_group.Add(1)
		go FetchData(&artist, &wait_group, api)
		wait_group.Wait()

		if artist.Id == 0 && artist.Name == "" {
			renderTemplate(w, "error.html", http.StatusNotFound, http.StatusNotFound)
			return
		}
		_, ok := artist.Locations.(string)
		if ok {
			wait_group.Add(3)
			go FetchData(&location, &wait_group, artist.Locations.(string))
			go FetchData(&relation, &wait_group, artist.Relations.(string))
			go FetchData(&dates, &wait_group, artist.ConcertDates.(string))
			wait_group.Wait()
			artist.ConcertDates = dates
			artist.Locations = location
			artist.Relations = relation
		}
		renderTemplate(w, "artist.html", artist, http.StatusOK)
	}
}

func (A ArtistHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && artistUrls.MatchString(r.URL.Path[1:]+"?"+r.URL.RawQuery):
		GetUser(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/" && len(r.URL.Query()) == 0:
		GetUsers(w,r)
	case r.Method == http.MethodGet && (r.URL.Path == "/artists/" || r.URL.Path == "/artists"  ) && len(r.URL.Query()) == 0:
		
		GetUsers(w, r)
	case r.Method == http.MethodPost:
		renderTemplate(w, "error.html", http.StatusMethodNotAllowed, http.StatusMethodNotAllowed)
	default:
		renderTemplate(w, "error.html", http.StatusNotFound, http.StatusNotFound)
	}
}
