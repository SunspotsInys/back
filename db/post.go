package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/SunspotsInys/thedoor/models"
	"github.com/SunspotsInys/thedoor/utils"
)

func GetPostDetail(p *models.PostWithTag, id uint64, isAdmin bool) error {
	if p == nil {
		return errors.New("can not pass in a nil value")
	}
	sqlStr := "SELECT * FROM `posts` WHERE `id` = ? %s LIMIT 1"
	if isAdmin {
		sqlStr = fmt.Sprintf(sqlStr, "")
	} else {
		sqlStr = fmt.Sprintf(sqlStr, " AND `public` = true ")
	}
	err := db.Get(p, sqlStr, id)
	if err != nil {
		logger.Error().Msgf("failed to get post info, sqlStr = %s, err = %v", sqlStr, err)
		return err
	}
	err = db.Select(
		&(p.Tags),
		"SELECT `tags`.`id`, `tags`.`name` "+
			"FROM `tags` "+
			"INNER JOIN `taglists` "+
			"ON `tags`.`id` = `taglists`.`tid` "+
			"WHERE `taglists`.`pid` = ?",
		p.ID,
	)
	if err != nil {
		logger.Error().Msgf("failed to get post tag info, err = %v", err)
		return err
	}
	return err
}

func GetPostList(ps *[]models.PostWithTag, page, len int, onlyPublic bool) error {
	sqlStr := " SELECT `id`, `title`, LEFT(`content`, 150) `content`, `createtime`, `public`, `top` " +
		" FROM `posts` " +
		" %s " +
		" ORDER BY `top` DESC, `id` DESC " +
		" LIMIT ? " +
		" OFFSET ? "
	if onlyPublic {
		sqlStr = fmt.Sprintf(sqlStr, " WHERE `public` = true ")
	} else {
		sqlStr = fmt.Sprintf(sqlStr, "")
	}
	log.Println(sqlStr)
	err := db.Select(ps, sqlStr, len, (page-1)*len)
	if err != nil {
		return err
	}

	for i := range *ps {
		err = db.Select(
			&((*ps)[i].Tags),
			"SELECT `tags`.`id`, `tags`.`name` "+
				"FROM `tags` "+
				"INNER JOIN `taglists` "+
				"ON `tags`.`id` = `taglists`.`tid` "+
				"WHERE `taglists`.`pid` = ?",
			(*ps)[i].ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPostTotal(onlyPublic bool) (int64, error) {
	sqlStr := "SELECT COUNT(*) FROM `posts` %s"
	if onlyPublic {
		sqlStr = fmt.Sprintf(sqlStr, " WHERE `public` = true ")
	} else {
		sqlStr = fmt.Sprintf(sqlStr, "")
	}
	var tot int64
	err := db.Get(&tot, sqlStr)
	return tot, err
}

func InsertPost(p *models.Post, tags *[]models.Tag) error {
	if p == nil {
		return errors.New("")
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
			err = errors.New("panic")
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	sf := utils.GetSnowflakeInstance()
	for i := range *tags {
		if (*tags)[i].ID == 0 {
			(*tags)[i].ID = sf.GetVal()
			rs, err := tx.Exec("INSERT INTO `tags` (`id`, `name`) VALUES(?, ?)", (*tags)[i].ID, (*tags)[i].Name)
			if err != nil {
				return err
			}
			n, err := rs.RowsAffected()
			if err != nil {
				return err
			}
			if n != 1 {
				return errors.New("exec insert error")
			}
		}
	}
	p.ID = sf.GetVal()
	rs, err := tx.Exec(
		"INSERT INTO `posts` (`id`, `title`, `content`, `createtime`, `public`, `top`) "+
			"VALUES (?, ?, ?, ?, ?, ?);",
		p.ID, p.Title, p.Content, p.CreateTime, p.Public, p.Top,
	)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec insert error")
	}

	stmt, err := tx.Prepare("INSERT INTO `thedoor`.`taglists` (`id`, `pid`, `tid`) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range *tags {
		rs, err = stmt.Exec(sf.GetVal(), p.ID, v.ID)
		if err != nil {
			return err
		}
		n, err := rs.RowsAffected()
		if err != nil {
			return err
		}
		if n != 1 {
			return errors.New("exec insert error")
		}
	}

	return nil
}

func GetPostSimpleyList(ps *[]models.PostSimplicity, start, len int) error {
	return db.Select(
		ps,
		"SELECT `id`, `title`, `createtime`"+
			" FROM `posts` "+
			" ORDER BY `id` DESC "+
			" LIMIT ? "+
			" OFFSET ? ",
		len,
		start,
	)
}
