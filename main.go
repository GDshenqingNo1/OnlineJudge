package main

import "OnlineJudge/boot"

func main() {

	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
	boot.SeverSetup()
}
