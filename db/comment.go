package db

import (
	"errors"

	"github.com/SunspotsInys/thedoor/models"
	"github.com/SunspotsInys/thedoor/utils"
)

func InsertComment(c *models.Comment) error {
	cid := utils.GetSnowflakeInstance()
	c.ID = cid.GetVal()
	rs, err := db.Exec(
		"INSERT INTO `comments`"+
			" (`id`, `pid`, `fid`, `content`, `createTime`, `cname`, `cemail`, `csite`) "+
			" VALUES(?, ?, ?, ?, ?, ?, ?, ?) ",
		c.ID, c.PID, c.FID, c.Content, c.CreateTime, c.CName, c.CEMail, c.CSite,
	)
	if err != nil {
		return err
	}
	num, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("inserted wrongly")
	}
	num, err = rs.LastInsertId()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("unpredictable data inserted")
	}
	return nil
}

func GetCommentsList(cs *[]models.Comments, pid uint64) error {
	err := db.Select(cs, "SELECT FROM `comments` WHERE `pid` = ? AND `fid` = \"\"", pid)
	if err != nil {
		return err
	}
	for i := 0; i < len(*cs); i++ {
		err = db.Select(
			&((*cs)[i].ChildComment),
			"SELECT FROM `comments` WHERE `pid` = ? AND `fid` = ?",
			pid, (*cs)[i].ID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
