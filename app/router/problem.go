package router

import (
	"OnlineJudge/app/api"
	"github.com/gin-gonic/gin"
)

type ProblemRouter struct{}

func (r *ProblemRouter) InitProblemRouter(router *gin.RouterGroup) gin.IRouter {
	ProblemRouter := router.Group("/problem")
	problemApi := api.Problem()
	{
		ProblemRouter.GET("info", problemApi.Problem().ProblemInfo)
		ProblemRouter.POST("/submit", problemApi.Problem().SubmitProblem)
		ProblemRouter.GET("/list", problemApi.Problem().ShowList)
		ProblemRouter.POST("/modify", problemApi.Problem().ModifyProblem)
	}
	return ProblemRouter
}
