package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/SunspotsInys/thedoor/logs"
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
		logs.Errorf("failed to get post info, sqlStr = %s, err = %v", sqlStr, err)
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
		logs.Errorf("failed to get post tag info, err = %v", err)
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

func GetPostListByTID(tags *[]models.PostWithSameTID, id uint64, admin bool) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	sqlStr := " SELECT `posts`.`id`, `posts`.`title`, `posts`.`createtime`, `posts`.`public` " +
		" FROM `taglists` INNER JOIN `posts` ON `posts`.`id` = `taglists`.`pid` " +
		" WHERE `taglists`.`tid` = ? %s" +
		" ORDER BY `posts`.`createtime` DESC "
	if admin {
		sqlStr = fmt.Sprintf(sqlStr, "")
	} else {
		sqlStr = fmt.Sprintf(sqlStr, " AND `posts`.public = 1")
	}
	return db.Select(tags, sqlStr, id)
}

func GetAchieve(tags *[]models.PostWithSameTID, admin bool) error {
	if tags == nil {
		return errors.New("can not pass in a nil value")
	}
	sqlStr := " SELECT `posts`.`id`, `posts`.`title`, `posts`.`createtime`, `posts`.`public` " +
		" FROM `posts` %s ORDER BY `posts`.`createtime` DESC  "
	if admin {
		sqlStr = fmt.Sprintf(sqlStr, "")
	} else {
		sqlStr = fmt.Sprintf(sqlStr, " WHERE `posts`.public = 1")
	}
	return db.Select(tags, sqlStr)
}

func UpdatePost(p *models.Post, tags *[]models.Tag) error {
	if p == nil {
		return errors.New("nil post")
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

	var ids = []uint64{}
	mp := map[uint64]bool{}
	err = tx.Select(&ids, "SELECT `tid` FROM `taglists` WHERE `pid` = ?", p.ID)
	if err != nil {
		return err
	}
	for _, v := range ids {
		mp[v] = false
	}

	log.Println(len(*tags))
	sf := utils.GetSnowflakeInstance()
	for i := range *tags {
		if (*tags)[i].ID == 0 {
			(*tags)[i].ID = sf.GetVal()
			_, err = tx.Exec("INSERT INTO `tags` (`id`, `name`) VALUES(?, ?)",
				(*tags)[i].ID, (*tags)[i].Name)
			if err != nil {
				return err
			}
			taglistid := sf.GetVal()
			_, err = tx.Exec("INSERT INTO `taglists` (`id`, `tid`, `pid`) VALUES(?, ?, ?)",
				taglistid, (*tags)[i].ID, p.ID)
			if err != nil {
				return err
			}
		} else {
			_, ok := mp[(*tags)[i].ID]
			if !ok {
				taglistid := sf.GetVal()
				_, err = tx.Exec("INSERT INTO `taglists` (`id`, `tid`, `pid`) VALUES(?, ?, ?)",
					taglistid, (*tags)[i].ID, p.ID)
				if err != nil {
					return err
				}
			} else {
				mp[(*tags)[i].ID] = true
			}
		}
	}

	for k, v := range mp {
		if !v {
			_, err = tx.Exec("DELETE FROM `tagslists` WHERE `pid` = ? AND `tid` = ?", p.ID, k)
			if err != nil {
				return err
			}
		}
	}

	_, err = tx.Exec(
		"UPDATE `posts` SET `title` = ?, `content` = ?, `public` = ?, `top` = ? WHERE `id` = ?",
		p.Title, p.Content, p.Public, p.Top, p.ID,
	)
	if err != nil {
		return err
	}
	log.Println("no error")

	return nil
}

func DeletePost(id uint64) error {

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

	_, err = tx.Exec("DELETE FROM `posts` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM `taglists` WHERE `pid` = ?", id)
	if err != nil {
		return err
	}

	return nil
}
