package service

import (
	"OnlineJudge/app/internal/service/internal/problem"
	"OnlineJudge/app/internal/service/internal/submission"
	"OnlineJudge/app/internal/service/internal/user"
)

var (
	insSubmission = submission.Group{}
	insUser       = user.Group{}
	insProblem    = problem.Group{}
)

func User() *user.Group {
	return &insUser
}

func Submission() *submission.Group {
	return &insSubmission
}

func Problem() *problem.Group {
	return &insProblem
}
