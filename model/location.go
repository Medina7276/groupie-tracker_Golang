package model

type Location struct {
	Id uint64 `json:"id"`
	Locations []string `json:"locations"`
	DatesUrl string `json:"dates"`
}
