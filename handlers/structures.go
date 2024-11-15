package handlers

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    any      `json:"locations"`
	ConcertDates any      `json:"concertDates"`
	Relations    any      `json:"relations"`
}

type Location struct {
	Id        int           `json:"id"`
	Locations []interface{} `json:"locations"`
	Dates     string        `json:"dates"`
}

type Relation struct {
	Id             int                    `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}

type Date struct {
	Id    int           `json:"id"`
	Dates []interface{} `json:"dates"`
}
