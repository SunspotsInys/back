package db

import (
	"errors"

	"github.com/SunspotsInys/thedoor/models"
)

func GetTagList(tags *[]models.Tags) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	return db.Select(tags, " SELECT `tags`.`id`, `tags`.`name`, COUNT(`taglists`.`tid`) num "+
		" FROM `taglists` "+
		" INNER JOIN posts "+
		" ON taglists.pid = posts.id "+
		" INNER JOIN `tags` "+
		" ON `taglists`.`tid` = `tags`.`id` "+
		" GROUP BY `taglists`.`tid` ",
	)
}

func GetTagListByID(tags *[]models.Tag, id int64) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	return db.Select(tags, "", id)
}
