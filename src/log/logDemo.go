package main

import (
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.LogPathSet("/Users/zhouzhenyong/tem/learn-go/logs/learn-go")

	serviceLog := log.GetLogger("test")
	serviceLog.SetLevel(logrus.DebugLevel)
	serviceLog.WithField("name", "zhou").Debug("haod")
	serviceLog.WithField("name", "zhou").Info("haod")

	time.Sleep(10000 * time.Second)
}
