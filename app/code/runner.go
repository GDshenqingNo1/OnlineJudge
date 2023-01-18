package code

import (
	g "OnlineJudge/app/global"
	"bytes"
	"go.uber.org/zap"
	"io"
	"os/exec"
)

func Run(uuid, testIn string) string {
	var path string
	path = "/home/code/OnlineJudge/app/code/" + uuid + "/main.go"
	cmd := exec.Command("go", "run", path)
	var stdErr, stdOut bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut

	stdIn, err := cmd.StdinPipe()
	if err != nil {
		g.Logger.Fatal("test in error.", zap.Error(err))
	}
	go func() {
		defer stdIn.Close()
		_, err = io.WriteString(stdIn, testIn)
		if err != nil {
			g.Logger.Fatal("test in error")
		}
	}()
	err = cmd.Run()
	if err != nil {
		g.Logger.Error("code build error", zap.Error(err))
		return stdErr.String()
	}
	return stdOut.String()

}
