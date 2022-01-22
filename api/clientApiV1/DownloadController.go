package clientApiV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLinks(c *gin.Context) {
	c.String(http.StatusOK, "GetLinks")
}

func GetPreview(c *gin.Context) {
	c.String(http.StatusOK, "GetLinks")
}

func GetArchive(c *gin.Context) {
	c.String(http.StatusOK, "GetLinks")
}
