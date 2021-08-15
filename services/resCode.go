package services

import "net/http"

type resCode uint64

// 2000+ 正常响应
const (
	codeSuccess resCode = iota + 2000
)

// 4000+ 请求数据异常
const (
	codeParamError resCode = 4000 + iota
	codePayloadError
	codeUsernameOrPasswordError
	codeNoRight
)

// 5000+ 响应异常
const (
	codeRefuse resCode = 5000 + iota
	codeServiceBusy
)

func (rc resCode) Msg() string {
	switch rc {
	case codeSuccess:
		return "成功访问"
	case codeParamError:
		return "参数错误"
	case codePayloadError:
		return "请求数据异常"
	case codeUsernameOrPasswordError:
		return "用户名或用户密码错误"
	case codeNoRight:
		return "僭越！！！"
	case codeRefuse:
		return "拒绝服务"
	case codeServiceBusy:
		return "服务繁忙"
	default:
		return ""
	}
}

func (rc resCode) StatusCode() int {
	switch rc {
	case codeSuccess:
		return http.StatusOK
	case codeParamError:
		return http.StatusBadRequest
	case codePayloadError:
		return http.StatusNotAcceptable
	case codeUsernameOrPasswordError:
		return http.StatusUnprocessableEntity
	case codeNoRight:
		return http.StatusForbidden
	case codeRefuse:
		return http.StatusInternalServerError
	case codeServiceBusy:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
