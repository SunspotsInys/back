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
		PID     uint64 `json:"pid"`
		FID     uint64 `json:"fid"`
		Content string `json:"content"`
		CName   string `json:"cname"`
		CEMail  string `json:"cemail"`
		CSite   string `json:"csite"`
	}{}

	// 解析请求数据
	err := c.BindJSON(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	if data.PID == 0 || data.Content == "" || data.CName == "" || data.CEMail == "" || data.CSite == "" {
		logger.Error().Msgf("the false data of comment, data = %+v", data)
		responseError(c, codeParamError)
		return
	}
	err = db.InsertComment(&models.Comment{
		PID:        data.PID,
		FID:        data.FID,
		Content:    data.Content,
		CreateTime: time.Now(),
		CName:      data.CName,
		CEMail:     data.CEMail,
		CSite:      data.CSite,
	})
	if err != nil {
		logger.Error().Msgf("Failed to create a comment, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, nil)
}
