package services

import (
	"log"
	"time"

	"github.com/SunspotsInys/thedoor/db"
	"github.com/SunspotsInys/thedoor/models"
	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {

	data := struct {
		Page int `form:"page"`
		Len  int `form:"len"`
	}{}
	// 解析请求数据
	err := c.BindQuery(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	logger.Info().Msgf("%+v", data)

	// 判断是否是管理员
	isAdmin := c.GetBool("isAdmin")
	log.Println("isAdmin = ", isAdmin)

	// 获取博客信息
	var ps []models.PostWithTag
	err = db.GetPostList(&ps, data.Page, data.Len, !isAdmin)
	if err != nil {
		logger.Error().Msgf("获取博客信息失败, err = %v", err.Error())
		responseError(c, codeServiceBusy)
		return
	}

	// 成功响应
	logger.Debug().Msgf("%+v", ps)
	responseSuccess(c, ps)
}

func GetPostTotal(c *gin.Context) {

	// 判断是否是管理员
	isAdmin := c.GetBool("admin")

	// 获取博客条数
	tot, err := db.GetPostTotal(!isAdmin)
	if err != nil {
		logger.Error().Msgf("获取博文条数失败, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}

	// 成功响应
	logger.Debug().Msgf("共%+v条博文", tot)
	responseSuccess(c, tot)
}

func GetPostDetail(c *gin.Context) {
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

	// 判断是否是管理员
	isAdmin := c.GetBool("isAdmin")

	var p models.PostWithTag
	err = db.GetPostDetail(&p, data.ID, isAdmin)
	if err != nil {
		logger.Error().Msgf("获取博文信息失败, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}

	// 成功响应
	logger.Debug().Msgf("%+v", p)
	responseSuccess(c, p)
}

func NewPost(c *gin.Context) {
	data := struct {
		Title   string       `json:"title"`
		Content string       `json:"content"`
		Public  bool         `json:"public"`
		Top     bool         `json:"top"` // 是否置顶
		Tags    []models.Tag `json:"tags"`
	}{}

	// 解析请求数据
	err := c.Bind(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}

	// 校验数据
	if data.Content == "" || data.Title == "" {
		logger.Error().Msgf("数据不规范, %+v", data)
		responseError(c, codeParamError)
	}

	p := models.Post{
		ID:         0,
		Title:      data.Title,
		Content:    data.Content,
		CreateTime: time.Now(),
		Public:     data.Public,
	}
	if data.Top {
		p.Top = time.Now().Unix()
	}
	err = db.InsertPost(&p, &data.Tags)
	if err != nil {
		logger.Error().Msgf("Faild to create a post, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, nil)
}

func GetPostSimpleList(c *gin.Context) {
	data := struct {
		Page int `form:"page"`
		Len  int `form:"len"`
	}{}

	// 解析请求数据
	err := c.BindQuery(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}

	// 查询数据库
	var ps []models.PostSimplicity
	err = db.GetPostSimpleyList(&ps, (data.Page-1)*data.Len, data.Len)
	if err != nil {
		logger.Error().Msgf("db error, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}

	responseSuccess(c, ps)
}

func GetAchieve(c *gin.Context) {
	var err error
	p := []models.PostWithSameTID{}
	err = db.GetAchieve(&p, c.GetBool("isAdmin"))
	if err != nil {
		logger.Error().Msgf("failed to get post list by tag id , err = %s", err.Error())
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, &p)
}
