package db

import (
	"errors"

	"github.com/SunspotsInys/thedoor/models"
)

func InsertComment(c *models.Comment) error {
	rs, err := db.Exec(
		"INSERT INTO `comments`"+
			" (`id`, `pid`, `fid`, `content`, `createtime`, `name`, `email`, `site`) "+
			" VALUES(?, ?, ?, ?, ?, ?, ?, ?) ",
		c.ID, c.PID, c.FID, c.Content, c.CreateTime, c.Name, c.EMail, c.Site,
	)
	if err != nil {
		return err
	}
	num, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("unpredictable data inserted")
	}
	return nil
}

func GetCommentsList(cs *[]models.Comments, pid uint64) error {
	err := db.Select(cs, " SELECT "+
		" `id`, `content`, `createtime`, `name`, `email`, `site` "+
		" FROM `comments` WHERE `pid` = ? AND `fid` = 0 ",
		pid,
	)
	if err != nil {
		return err
	}
	for i := 0; i < len(*cs); i++ {
		err = db.Select(&((*cs)[i].Children), " SELECT "+
			" `id`, `fid`, `content`, `createtime`, `name`, `email`, `site` "+
			" FROM `comments` "+
			" WHERE `pid` = ? AND `fid` = ? "+
			" ORDER BY `id` ",
			pid, (*cs)[i].ID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
