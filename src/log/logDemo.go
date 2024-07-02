package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {

	serviceLog := log.GetLogger("test")
	serviceLog.SetLevel(logrus.DebugLevel)
	serviceLog.WithField("name", "zhou").Debug("haod")
	serviceLog.WithField("name", "zhou").Info("haod")

	time.Sleep(10000 * time.Second)
}
