package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var serviceLog = logrus.New()
var path = "/Users/zhouzhenyong/tem/learn-go/logs/access1"
var infoRotate, _ = rotatelogs.New(path+"-info.log.%Y%m%d", rotatelogs.WithLinkName(path+"-info.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var warnRotate, _ = rotatelogs.New(path+"-warn.log.%Y%m%d", rotatelogs.WithLinkName(path+"-warn.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var errorRotate, _ = rotatelogs.New(path+"-error.log.%Y%m%d", rotatelogs.WithLinkName(path+"-error.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))

func main() {
	// INFO[2021-12-14 17:28:28]
	serviceLog.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.000", FullTimestamp: true}
	// 显示行号
	serviceLog.SetReportCaller(true)

	// 输出到标准输出，而不是默认的标准错误
	// 可以是任何io.Writer，请参阅下面的文件例如日志。
	serviceLog.Out = os.Stdout
	//
	//// 仅记录严重警告以上。
	//serviceLog.Level = logrus.InfoLevel
	//
	//serviceLog.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//}).Info("A walrus appears")
	//
	//serviceLog.Info("haode")

	//	`WithLinkName` 为最新的日志建立软连接
	//	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	//	 WithMaxAge 和 WithRotationCount二者只能设置一个
	//	  `WithMaxAge` 设置文件清理前的最长保存时间
	//	  `WithRotationCount` 设置文件清理前最多保存的个数

	//serviceLog.AddHook(&DefaultFieldHook{})
	serviceLog.Debug("debug-data")
	serviceLog.Info("info-data")
	serviceLog.Warn("warn-data")
	serviceLog.Error("error-data")
	//
	//serviceLog.SetOutput(rl)
}

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		entry.Logger.Out = errorRotate
	case logrus.WarnLevel:
		entry.Logger.Out = warnRotate
	case logrus.InfoLevel:
		entry.Logger.Out = infoRotate
	case logrus.DebugLevel:
		entry.Logger.Out = infoRotate
	default:
		return nil
	}
	return nil
}

func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
