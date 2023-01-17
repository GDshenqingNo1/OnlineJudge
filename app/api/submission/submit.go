package submission

import (
	"OnlineJudge/app/internal/model/resp"
	"OnlineJudge/app/internal/model/submission"
	"OnlineJudge/app/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type SubmitApi struct{}

var insSubmission SubmitApi

func (a *SubmitApi) Submit(c *gin.Context) {
	//fmt.Println(c.Request.Body)
	problemId, _ := strconv.Atoi(c.Query("problem_id"))
	username := c.Query("username")
	code, _ := io.ReadAll(c.Request.Body)
	submissionSubject := &submission.Submission{
		Uuid:      service.Submission().Submission().GetUuid(),
		Username:  username,
		ProblemId: problemId,
	}

	service.Submission().Submission().SaveCode(code, submissionSubject)
	err := service.Submission().Submission().SaveRecord(c, submissionSubject)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("save submission record failed:%v", err))
		return
	}
	problem, err := service.Problem().Problem().GetProblem(c, submissionSubject.ProblemId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, fmt.Sprintf("get problem from mysql failed,err:%v", err))
	}
	if out, correct := service.Submission().Submission().TestCode(problem, submissionSubject); correct {
		resp.ResponseSuccess(c, 200, "answer accept."+" "+out)
		return
	} else {
		resp.ResponseSuccess(c, 200, "wrong answer"+" "+out)
	}
}
