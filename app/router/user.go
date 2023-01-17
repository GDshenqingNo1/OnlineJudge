package router

import (
	"OnlineJudge/app/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	UserRouter := router.Group("/user")
	userApi := api.User()
	{
		UserRouter.POST("/register", userApi.Sign().Register)
		UserRouter.POST("/login", userApi.Sign().Login)
		UserRouter.POST("/sendCode", userApi.Sign().SendCode)
	}

	return UserRouter

}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	return userRouter
}
