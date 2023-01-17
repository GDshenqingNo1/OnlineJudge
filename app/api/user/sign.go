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
	username := c.PostForm("username")
	password := c.PostForm("password")
	mail := c.PostForm("mail")
	code := c.PostForm("code")
	if username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null.")
		return
	}
	if password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null.")
		return
	}
	err := service.User().User().CheckUserIsExist(c, username)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err.")
		} else if err.Error() == "username already exist" {
			resp.ResponseSuccess(c, http.StatusOK, "username already exist.")
		}
		return
	}
	err = service.User().User().CheckCode(c, mail, code)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
			return
		} else if err.Error() == "incorrect code" {
			resp.ResponseFail(c, http.StatusBadRequest, "incorrect code")
			return
		}
	}
	userSubject := &user.User{}
	encryptedPassword := service.User().User().EncryptPassword(password)
	userSubject.Username = username
	userSubject.Password = encryptedPassword
	userSubject.Mail = mail
	err = service.User().User().CreateUser(c, userSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("register failed：%v", err))
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "register successfully")
}

func (a *SignApi) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null.")
		return
	}
	if password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null.")
		return
	}
	userSubject := &user.User{
		Username: username,
		Password: service.User().User().EncryptPassword(password),
	}
	err := service.User().User().CheckPassword(c, userSubject)
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
			resp.ResponseFail(c, http.StatusBadRequest, "mail is already signed")
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
