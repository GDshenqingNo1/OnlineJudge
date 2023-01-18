package user

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/model/resp"
	"OnlineJudge/app/internal/model/user"
	"OnlineJudge/app/internal/service"
	"OnlineJudge/utils/cookie"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignApi struct{}

var insSign = SignApi{}

func (a *SignApi) Register(c *gin.Context) {
	var userSubject = &user.User{}
	err := c.BindJSON(&userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, fmt.Sprintf("request json form error:%v", err))
		return
	}
	if userSubject.Username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null.")
		return
	}
	if userSubject.Password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null.")
		return
	}
	err = service.User().User().CheckUserIsExist(c, userSubject.Username)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err.")
		} else if err.Error() == "username already exist" {
			resp.ResponseSuccess(c, http.StatusOK, "username already exist.")
		}
		return
	}
	err = service.User().User().CheckCode(c, userSubject.Mail, userSubject.Code)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
			return
		} else if err.Error() == "incorrect code" {
			resp.ResponseFail(c, http.StatusOK, "incorrect code")
			return
		}
	}
	encryptedPassword := service.User().User().EncryptPassword(userSubject.Password)
	userSubject.Password = encryptedPassword
	err = service.User().User().CreateUser(c, userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("register failed：%v", err))
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "register successfully")
}

func (a *SignApi) Login(c *gin.Context) {
	var userSubject = &user.User{}
	err := c.BindJSON(&userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusOK, fmt.Sprintf("request json form error:%v", err))
		return
	}
	if userSubject.Username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null.")
		return
	}
	if userSubject.Password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null.")
		return
	}
	userSubject.Password = service.User().User().EncryptPassword(userSubject.Password)
	err = service.User().User().CheckPassword(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			resp.ResponseFail(c, http.StatusInternalServerError, err.Error())
		case "invalid username or password":
			resp.ResponseFail(c, http.StatusBadRequest, err.Error())
		}
		return
	}
	tokenString, err := service.User().User().GenerateToken(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
		}
	}
	cookieConfig := g.Config.App.Cookie

	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
		cookie.Option{
			Config: cookieConfig.Cookie,
			Ctx:    c,
		})
	cookieWriter.Set("x-token", tokenString)

	resp.ResponseSuccess(c, http.StatusOK, "login successfully")
}
func (a *SignApi) ResetPassword(c *gin.HandlerFunc) {

}

func (a *SignApi) SendCode(c *gin.Context) {
	mail := c.PostForm("mail")
	if mail == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "mail cna not be null.")
	}
	err := service.User().User().CheckMailIsExist(c, mail)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("check mail failed ,err:%v", err))
			return
		} else if err.Error() == "mail is already signed" {
			resp.ResponseFail(c, http.StatusOK, "mail is already signed")
			return
		}
	}
	err = service.User().User().SendCode(c, mail)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("send code error:%v", err))
		return
	} else {
		resp.ResponseSuccess(c, http.StatusOK, "send code successfully.")
	}
}
