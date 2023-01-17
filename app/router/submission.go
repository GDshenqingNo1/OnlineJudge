package router

import (
	"OnlineJudge/app/api"
	"github.com/gin-gonic/gin"
)

type SubmissionRouter struct{}

func (r *SubmissionRouter) InitSubmissionRouter(router *gin.RouterGroup) gin.IRouter {
	submissionApi := api.Submission()
	{
		router.POST("/submit", submissionApi.Submission().Submit)
	}
	return router
}
