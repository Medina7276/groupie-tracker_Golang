package model

type Artist struct {
	Id int64 `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate uint16 `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`

	LocationsUrl string `json:"locations"`
	//Locations []Location

	ConcertDatesUrl string `json:"concertDates"`
	//ConcertDates []Date

	RelationsUrl string `json:"relations"`
	//Relations []Relation
}
