package main

import (
	"github.com/simonalong/tools"
	"github.com/sirupsen/logrus"
)

func main() {
	tools.LogPathSet("/Users/zhouzhenyong/tem/learn-go/logs/learn-go")

	serviceLog := tools.GetLogger("test")
	serviceLog.SetLevel(logrus.DebugLevel)
	serviceLog.WithField("name", "zhou").Debug("haod")
	serviceLog.WithField("name", "zhou").Info("haod")
}
