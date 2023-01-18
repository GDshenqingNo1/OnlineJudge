package router

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true), middleware.CORS())

	routerGroup := new(Group)
	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(publicGroup)
		routerGroup.InitProblemRouter(publicGroup)
		routerGroup.InitSubmissionRouter(publicGroup)
	}
	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{

	}
	g.Logger.Info("initialize routers successfully")
	return r
}
