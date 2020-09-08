package repository

import (
	"encoding/json"
	"groupie-tracker/model"
)


func GetLocationsFromArtistAsync(url string, chanLoc chan<- model.Location) {
	location := model.Location{}

	_ = getLocation(url, &location)

	chanLoc <- location
}

func GetLocationsFromArtist(url string) (model.Location, error) {
	location := model.Location{}

	err := getLocation(url, &location)
	if err != nil {
		return location, err
	}

	return location, nil
}

func getLocation(url string, target *model.Location) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}