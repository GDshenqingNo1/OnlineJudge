package problem

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/model/problem"
	"OnlineJudge/app/internal/model/resp"
	"OnlineJudge/app/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ProblemApi struct{}

var insProblem ProblemApi

func (a *ProblemApi) ShowList(c *gin.Context) {
	//pageStr := c.Query("page")
	//pageSize, _ := strconv.Atoi(pageStr)
	list, err := service.Problem().Problem().GetProblemList(c)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("get problem list error:%v", err))
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"mag":  "get problem list successfully",
		"data": list,
	})
}

func (a *ProblemApi) SubmitProblem(c *gin.Context) {
	var problemSubject = &problem.Problem{}
	err := c.BindJSON(&problemSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, fmt.Sprintf("request JSON form error:%v", err))
		return
	}
	err = service.Problem().Problem().SaveProblem(c, problemSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("submit problem failed err:%v", err))
		return
	}
	resp.ResponseSuccess(c, 200, "submit problem successfully.")
}

func (a *ProblemApi) ModifyProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := service.Problem().Problem().CheckProblemDoesExist(c, id)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
			return
		} else if err.Error() == "problem is not exist" {
			resp.ResponseFail(c, http.StatusOK, "problem is not exist")
			return
		}
	}
	var problemSubject = &problem.Problem{}

	err = c.BindJSON(&problemSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, fmt.Sprintf("request json error:%v", err))
		return
	}
	problemSubject.Id = id
	err = service.Problem().Problem().UpdateProblemInfo(c, problemSubject)
	if err != nil {
		g.Logger.Error("update problem info err", zap.Error(err))
		resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "modify problem successfully")
}

func (a *ProblemApi) ProblemInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := service.Problem().Problem().CheckProblemDoesExist(c, id)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
			return
		} else if err.Error() == "problem is not exist" {
			resp.ResponseFail(c, http.StatusOK, "problem is not exist")
			return
		}
	}
	problemSubject, err := service.Problem().Problem().GetProblem(c, id)
	if err != nil {
		if err.Error() == "internal err" {
			resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("internal err:%v", err))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "get problem information successfully",
		"data": problemSubject,
	})
}
