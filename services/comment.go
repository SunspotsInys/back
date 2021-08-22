package services

import (
	"time"

	"github.com/SunspotsInys/thedoor/db"
	"github.com/SunspotsInys/thedoor/models"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	data := struct {
		ID uint64 `form:"pid"`
	}{}

	// 解析请求数据
	err := c.BindQuery(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	var cs []models.Comments
	err = db.GetCommentsList(&cs, data.ID)
	if err != nil {
		logger.Error().Msgf("获取评论失败, uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	responseSuccess(c, &cs)
}

func NewComments(c *gin.Context) {
	data := struct {
		PID     uint64 `json:"pid,string"`
		FID     uint64 `json:"fid,string"`
		Content string `json:"content"`
		Name    string `json:"name"`
		EMail   string `json:"email"`
		Site    string `json:"site"`
	}{}

	// 解析请求数据
	err := c.BindJSON(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	if data.PID == 0 || data.Content == "" || data.Name == "" || data.EMail == "" || data.Site == "" {
		logger.Error().Msgf("the false data of comment, data = %+v", data)
		responseError(c, codeParamError)
		return
	}
	id := idGen.GetVal()
	err = db.InsertComment(&models.Comment{
		ID:         id,
		PID:        data.PID,
		FID:        data.FID,
		Content:    data.Content,
		CreateTime: time.Now(),
		Name:       data.Name,
		EMail:      data.EMail,
		Site:       data.Site,
	})
	if err != nil {
		logger.Error().Msgf("Failed to create a comment, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, id)
}
