package model

type Index struct {
	Relations []Relation `json:"index"`
}

type Relation struct {
	Id uint64 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
