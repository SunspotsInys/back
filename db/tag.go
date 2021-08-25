package db

import (
	"errors"

	"github.com/SunspotsInys/thedoor/models"
)

func GetTagInfoList(tags *[]models.Tags) error {
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

func GetTagList(ts *[]models.Tag) error {
	if ts == nil {
		return errors.New("can not pass in a nil value")
	}
	return db.Select(ts, "SELECT * FROM `tags`")
}

func GetTagInfo(t *models.Tag, id uint64) error {
	return db.Get(t, "SELECT `name` FROM `tags` WHERE `id` = ? LIMIT 1 ", id)
}
