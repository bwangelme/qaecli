package main

import (
	"qaecli/cmd"

	"github.com/sirupsen/logrus"
)

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	initLog()
	cmd.Execute()
}
