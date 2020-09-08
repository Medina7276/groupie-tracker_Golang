package dto

import "groupie-tracker/model"

type Artist struct {
	Id int64 `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate uint16 `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`

	Location model.Location `json:"locations"`

	ConcertDates model.Date `json:"concertDates"`

	Relations model.Relation `json:"relations"`
}
