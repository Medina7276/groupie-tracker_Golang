package repository

import "groupie-tracker/model"

func GetRelationsFromArtist(url string) (model.Relation, error) {
	var relations model.Relation

	err := get(url, &relations)
	if err != nil {
		return relations, err
	}

	return relations, nil
}

func GetRelationsFromArtistAsync(url string, chanLoc chan<- model.Relation) {
	relation := model.Relation{}

	_ = get(url, &relation)

	chanLoc <- relation
}