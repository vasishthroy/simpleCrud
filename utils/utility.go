package utils

import (
	"math/rand"
	"simplecrud/model"
	"strconv"
)

func CheckID(id string, data *[]model.Anime) bool {
	for _, item := range *data {
		if item.ID == id {
			return true
		}
	}
	return false
}

func CreateID() string {
	return strconv.Itoa(rand.Intn(1000))
}

func DeleteRecord(id string, data *[]model.Anime) {

	// Check for the ID that needs to be deleted by looping through Animes slice
	for index, item := range *data {
		if item.ID == id {

			// Delete the record from the slice
			*data = append((*data)[:index], (*data)[index+1:]...)

			return
		}
	}
}
