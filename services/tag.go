package services

import (
	"github.com/SunspotsInys/thedoor/db"
	"github.com/SunspotsInys/thedoor/models"
	"github.com/gin-gonic/gin"
)

func GetTagList(c *gin.Context) {
	tags := new([]models.Tags)
	err := db.GetTagList(tags)
	if err != nil {
		logger.Error().Msgf("failed to get tag list, err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, tags)
}
