package repository

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/model"
	"net/http"
	"strconv"
	"time"
)

const artistsUrl string = "https://groupietrackers.herokuapp.com/api/artists"
var client = &http.Client{}

func GetArtists() ([]model.Artist, error) {
	var artists []model.Artist
	err := getbench(artistsUrl, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func GetArtistById(id int)  (*model.Artist, error) {
	artist := &model.Artist{}

	err := get(artistsUrl + "/" + strconv.Itoa(id), &artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func get(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func getbench(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()



	start := time.Now()

	err = json.NewDecoder(r.Body).Decode(target)

	elapsed := time.Since(start)
	fmt.Printf("all took %s \n", elapsed)

	if err != nil {
		return err
	}

	return nil
}
