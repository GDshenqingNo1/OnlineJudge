package api

import (
	"OnlineJudge/app/api/problem"
	"OnlineJudge/app/api/submission"
	"OnlineJudge/app/api/user"
)

var (
	insUser       = user.Group{}
	insProblem    = problem.Group{}
	insSubmission = submission.Group{}
)

func User() *user.Group {
	return &insUser
}
func Problem() *problem.Group {
	return &insProblem
}
func Submission() *submission.Group {
	return &insSubmission
}
