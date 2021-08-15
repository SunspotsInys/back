package services

import (
	"github.com/SunspotsInys/thedoor/configs"
	"github.com/SunspotsInys/thedoor/utils"
	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		logger.Error().Msgf("Failed to parse data , uri = %s, err = %v", c.Request.RequestURI, err)
		responseError(c, codePayloadError)
		return
	}
	if data.Username != configs.Conf.AdminUsername || data.Password != configs.Conf.AdminPassword {
		logger.Error().Msgf("Username or password is incorrect, username = %s, password = %s", data.Username, data.Password)
		responseError(c, codeUsernameOrPasswordError)
		return
	}
	token, err := utils.GenToken(data.Username)
	if err != nil {
		logger.Error().Msgf("Failed to generate token, err = %v", err)
		responseError(c, codeServiceBusy)
		return
	}
	responseSuccess(c, &struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Token    string `json:"token"`
	}{
		Username: configs.Conf.AdminUsername,
		Avatar:   configs.Conf.AdminAvatar,
		Token:    token,
	})
}

func ParseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug().Msgf("%+v", c.Request)
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			token = c.Query("token")
			logger.Debug().Msg(token)
		}
		uname := utils.ParseToken(token)
		logger.Debug().Msg(uname)
		if uname != "" {
			c.Set("isAdmin", true)
			c.Set("uname", uname)
		}
	}
}

func JudgeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, isExists := c.Get("isAdmin")
		if !isExists {
			logger.Error().Msgf("Fake Admin")
			c.Abort()
			responseError(c, codeNoRight)
		}
	}
}
