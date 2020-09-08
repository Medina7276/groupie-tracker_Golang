package repository

import "groupie-tracker/model"

func GetConcertDatesFromArtist(url string) (model.Date, error) {
	dates := model.Date{}

	err := get(url, &dates)
	if err != nil {
		return dates, err
	}

	return dates, nil
}

func GetConcertDatesFromArtistAsync(url string, chanLoc chan<- model.Date) {
	date := model.Date{}

	_ = get(url, date)

	chanLoc <- date
}