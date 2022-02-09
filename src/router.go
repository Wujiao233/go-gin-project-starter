package src

import (
	"git.caimi-inc.com/stanlee/approval-chain-go/src/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/ping", apis.Ping)

	//router.GET("/hivetest", hive.HdfsTest)
	return router
}
