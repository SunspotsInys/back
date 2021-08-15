package db

import (
	"errors"

	"github.com/SunspotsInys/thedoor/models"
)

func GetTagList(tags *[]models.Tag) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	return db.Select(tags, "SELECT * FROM `tags`")
}

func GetTagListByID(tags *[]models.Tag, id int64) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	return db.Select(tags, "", id)
}
