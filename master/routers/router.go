package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"../api"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//router.Use(gin.Logger())

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[Hound Master]%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	nodesGroup := router.Group("/nodes")
	{
		nodesGroup.POST("join",api.PostNodeJoin)
		nodesGroup.GET("list",api.GetNodeList)
	}
	tokensGroup := router.Group("/tokens")
	{
		tokensGroup.POST("send",api.SendToken)
		tokensGroup.GET("list",api.GetTokenList)
	}
	return router
}
