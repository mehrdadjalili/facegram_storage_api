package clientApiV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUpload(c *gin.Context) {
	c.String(http.StatusOK, "NewUpload")
}
