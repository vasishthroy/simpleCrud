package utils

import "simpleCrud/model"

func CheckID(id string, data []model.Anime) bool {
	for _, item := range data {
		if item.ID == id {
			return true
		}
	}
	return false
}
