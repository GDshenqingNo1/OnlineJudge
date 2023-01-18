package submission

import (
	"OnlineJudge/app/code"
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/model/problem"
	"OnlineJudge/app/internal/model/submission"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
)

type SSubmission struct{}

var insSubmission SSubmission

func (s *SSubmission) GetUuid() string {
	u := uuid.New()
	return u.String()
}

func (s *SSubmission) SaveCode(code []byte, submission *submission.Submission) {
	_ = os.Mkdir("/home/code/OnlineJudge/app/code/"+submission.Uuid, 0777)
	_, err := os.Create("/home/code/OnlineJudge/app/code/" + submission.Uuid + "/main.go")
	if err != nil {
		g.Logger.Fatal("save code to local error.", zap.Error(err))

	}
	file, _ := os.OpenFile("/home/code/OnlineJudge/app/code/"+submission.Uuid+"/main.go", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	defer file.Close()
	_, err = file.Write(code)
	if err != nil {
		g.Logger.Fatal("write code to file failed.")
	}
}

func (s *SSubmission) SaveRecord(ctx context.Context, submission *submission.Submission) error {
	err := g.MysqlDB.WithContext(ctx).
		Table("submission").
		Create(submission).Error
	return err
}

func (s *SSubmission) TestCode(problem *problem.Problem, submission *submission.Submission) (response string, correct bool) {
	out := code.Run(submission.Uuid, problem.TestIn)
	if out == problem.TestOut {
		return out, true
	}
	return out, false
}
