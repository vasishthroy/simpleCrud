package utils

import (
	"math/rand"
	"simplecrud/model"
	"strconv"
)

func CheckID(id string, data []model.Anime) bool {
	for _, item := range data {
		if item.ID == id {
			return true
		}
	}
	return false
}

func CreateID() string {
	return strconv.Itoa(rand.Intn(1000))
}
