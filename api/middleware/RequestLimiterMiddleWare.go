package middleware

import (
	"facegram_file_server/config"
	"facegram_file_server/model/responsemodel"
	"facegram_file_server/pkg/apiresponse"
	"facegram_file_server/pkg/requestlimiter"
	"github.com/gin-gonic/gin"
)

func AllRouteRequestLimiter(c *gin.Context) {
	ip := c.ClientIP()
	cfg, err := config.GetRequestLimiterConfigBuilder("ALL_ROUTE_REQUEST_LIMITER")
	if err != nil {
		apiresponse.SendInternalError(c)
	} else {
		result, t, tt, e0 := requestlimiter.CheckAndUpdate(ip, cfg)
		next(c, result, t, tt, e0)
	}
}

func next(c *gin.Context, result string, t int, tt string, e0 error) {
	if e0 != nil {
		apiresponse.SendInternalError(c)
	} else {
		if result == "open" {
			c.Next()
		} else if result == "close" {
			dataModel := responsemodel.RequestLimiterInfo{
				TimeType:  tt,
				LimitTime: t,
			}
			apiresponse.SendRequestLimit(c, dataModel)
		} else {
			apiresponse.SendInternalError(c)
		}
	}
}
