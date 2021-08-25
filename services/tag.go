package services

import (
	"strconv"

	"github.com/SunspotsInys/thedoor/db"
	"github.com/SunspotsInys/thedoor/models"
	"github.com/gin-gonic/gin"
)

func GetTagInofList(c *gin.Context) {
	tags := new([]models.Tags)
	err := db.GetTagInfoList(tags)
	if err != nil {
		logger.Error().Msgf("failed to get tag list, err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, tags)
}

func GetTagList(c *gin.Context) {
	tags := new([]models.Tag)
	err := db.GetTagList(tags)
	if err != nil {
		logger.Error().Msgf("failed to get tag list, err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, tags)
}

func GetPostByTag(c *gin.Context) {
	tid := c.Param("id")
	id, err := strconv.ParseUint(tid, 10, 64)
	if err != nil {
		logger.Error().Msgf("failed to get tag id, err = %s", err.Error())
		responseError(c, codeParamError)
		return
	}
	p := []models.PostWithSameTID{}
	err = db.GetPostListByTID(&p, id, c.GetBool("isAdmin"))
	if err != nil {
		logger.Error().Msgf("failed to get post list by tag id , err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	t := models.Tag{}
	err = db.GetTagInfo(&t, id)
	if err != nil {
		logger.Error().Msgf("failed to get post list by tag id , err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, &struct {
		Name string                    `json:"name"`
		Data *[]models.PostWithSameTID `json:"data"`
	}{
		Name: t.Name,
		Data: &p,
	})
}
