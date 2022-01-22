package apiresponse

import (
	"facegram_file_server/config"
	"facegram_file_server/model/responsemodel"
	"facegram_file_server/pkg/requestlimiter"

	"github.com/gin-gonic/gin"
)

func send(c *gin.Context, result string, msg string, statusCode int, data interface{}) {
	model := responsemodel.BaseStructRequest{
		Result:     result,
		Message:    msg,
		StatusCode: statusCode,
		Data:       data,
	}
	c.AbortWithStatusJSON(statusCode, model)
	return
}

//SendInternalError func
func SendInternalError(c *gin.Context) {
	send(c, ResultError, "error", StatusInternalError, nil)
}

//SendUnAuthorize func
func SendUnAuthorize(c *gin.Context) {
	send(c, ResultError, "invalid-token", StatusUnAuthorize, nil)
}

//SendInvalidInputError func
func SendInvalidInputError(c *gin.Context) {
	send(c, ResultError, "invalid-input", StatusInvalidInput, nil)
}

//SendError CustomizeAble func
func SendError(c *gin.Context, msg string, data interface{}) {
	send(c, ResultError, msg, StatusInternalError, data)
}

//SendSuccess CustomizeAble func
func SendSuccess(c *gin.Context, msg string, data interface{}) {
	send(c, ResultOk, msg, StatusOk, data)
}

//SendRequestLimit CustomizeAble func
func SendRequestLimit(c *gin.Context, data interface{}) {
	send(c, ResultError, "request-limit", StatusRequestLimit, data)
}

//SendSuccessAndUpdateRequestLimiter CustomizeAble func
func SendSuccessAndUpdateRequestLimiter(c *gin.Context, msg string, data interface{},
	id string, cfg *config.RequestLimiterItemDetail) {
	_, _ = requestlimiter.Set(id, cfg)
	send(c, ResultOk, msg, StatusOk, data)
}

//SendErrorAndUpdateRequestLimiter CustomizeAble func
func SendErrorAndUpdateRequestLimiter(c *gin.Context, msg string, data interface{},
	id string, cfg *config.RequestLimiterItemDetail) {
	_, _ = requestlimiter.Set(id, cfg)
	send(c, ResultError, msg, StatusInternalError, data)
}

//SendUnAuthorizeAndUpdateRequestLimiter CustomizeAble func
func SendUnAuthorizeAndUpdateRequestLimiter(c *gin.Context, id string,
	cfg *config.RequestLimiterItemDetail) {
	_, _ = requestlimiter.Set(id, cfg)
	send(c, ResultError, "invalid-token", StatusUnAuthorize, nil)
}

//SendForbidden CustomizeAble func
func SendForbidden(c *gin.Context, data interface{}) {
	send(c, ResultError, MessageForbidden, StatusForbidden, data)
}
