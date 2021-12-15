package main

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"sync"
	"time"
)

var serviceLog = logrus.New()
var path = "/Users/zhouzhenyong/tem/learn-go/logs/access1"
var infoRotate, _ = rotatelogs.New(path+"-info.log.%Y%m%d", rotatelogs.WithLinkName(path+"-info.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var warnRotate, _ = rotatelogs.New(path+"-warn.log.%Y%m%d", rotatelogs.WithLinkName(path+"-warn.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var errorRotate, _ = rotatelogs.New(path+"-error.log.%Y%m%d", rotatelogs.WithLinkName(path+"-error.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var fatalRotate, _ = rotatelogs.New(path+"-fatal.log.%Y%m%d", rotatelogs.WithLinkName(path+"-error.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
var panicRotate, _ = rotatelogs.New(path+"-panic.log.%Y%m%d", rotatelogs.WithLinkName(path+"-error.log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))

func main() {
	//// INFO[2021-12-14 17:28:28]
	//formater2 := &logrus.TextFormatter{
	//	//ForceColors: true,
	//	//DisableColors: false,
	//	////ForceQuote: true,
	//	//DisableQuote: true,
	//	//EnvironmentOverrideColors: true,
	//	////DisableTimestamp: true,
	//	//FullTimestamp: true,
	//	//TimestampFormat: "2006-01-02 15:04:05.000",
	//	//DisableSorting: true,
	//	//DisableLevelTruncation: true,
	//	//PadLevelText: true,
	//	//QuoteEmptyFields: true,
	//	//FieldMap: logrus.FieldMap{
	//	//	logrus.FieldKeyTime:  "[timestamp]",
	//	//	logrus.FieldKeyLevel: "[level]",
	//	//	logrus.FieldKeyMsg:   "[message]",
	//	//},
	//
	//	ForceQuote:true,    //键值对加引号
	//	TimestampFormat:"2006-01-02 15:04:05",  //时间格式
	//	FullTimestamp:true,
	//}
	//
	////serviceLog.Formatter = formater2
	//// 显示行号
	serviceLog.SetReportCaller(true)
	//
	//// 输出到标准输出，而不是默认的标准错误
	//// 可以是任何io.Writer，请参阅下面的文件例如日志。
	//serviceLog.Out = os.Stdout
	////
	////// 仅记录严重警告以上。
	////serviceLog.Level = logrus.InfoLevel
	////
	////serviceLog.WithFields(logrus.Fields{
	////	"animal": "walrus",
	////}).Info("A walrus appears")
	////
	////serviceLog.Info("haode")
	//
	////	`WithLinkName` 为最新的日志建立软连接
	////	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	////	 WithMaxAge 和 WithRotationCount二者只能设置一个
	////	  `WithMaxAge` 设置文件清理前的最长保存时间
	////	  `WithRotationCount` 设置文件清理前最多保存的个数
	//
	//serviceLog.SetFormatter(&MyFormatter{})
	//lfHook := lfshook.NewHook(lfshook.WriterMap{
	//	logrus.DebugLevel: infoRotate, // 为不同级别设置不同的输出目的
	//	logrus.InfoLevel:  infoRotate,
	//	logrus.WarnLevel:  infoRotate,
	//	logrus.ErrorLevel: infoRotate,
	//	logrus.FatalLevel: infoRotate,
	//	logrus.PanicLevel: infoRotate,
	//}, &MyFormatter{})
	////}, &logrus.JSONFormatter{})
	//serviceLog.AddHook(lfHook)
	//
	//serviceLog.Debug("debug-data111")
	//serviceLog.Info("info-data222")
	//serviceLog.Warn("warn-data333")
	//serviceLog.Error("error-data44")
	////
	////serviceLog.SetOutput(rl)
	//
	//
	//log.SetFormatter(&log.TextFormatter{})
	//switch level := serviceLog.Level; level {
	///*
	//   如果日志级别不是debug就不要打印日志到控制台了
	//*/
	//case logrus.DebugLevel:
	//	serviceLog.SetLevel(logrus.DebugLevel)
	//	serviceLog.SetOutput(os.Stderr)
	//case "info":
	//	serviceLog.SetLevel(logrus.InfoLevel)
	//case "warn":
	//	serviceLog.SetLevel(logrus.WarnLevel)
	//case "error":
	//	serviceLog.SetLevel(logrus.ErrorLevel)
	//default:
	//	serviceLog.SetLevel(logrus.InfoLevel)
	//}

	serviceLog.Formatter = &MyFormatter{}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: infoRotate, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  infoRotate,
		logrus.WarnLevel:  warnRotate,
		logrus.ErrorLevel: errorRotate,
		logrus.FatalLevel: fatalRotate,
		logrus.PanicLevel: panicRotate,
	}, &MyFormatter{})
	serviceLog.AddHook(lfHook)

	//serviceLog.WithField("name", "ball").WithField("say", "hi").Debug("info log")
	serviceLog.WithField("name", "ball").WithField("say", "hi").Info("我们大沙发了晶澳科技到地方了卡接到飞龙卡机的阿里看电视剧非了；奥科吉的")
	serviceLog.WithField("name", "ball").WithField("say", "hi").Warn("asdkfasdf ")
	serviceLog.WithField("name", "ball").WithField("say", "hi").Error("我们大沙发了晶澳科技到地方了卡接到飞龙卡机的阿里看电视剧非了；奥科吉的")
}

type MyFormatter struct {
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	entry.Caller = getCaller()

	var fields []string
	for k, v := range entry.Data {
		fields = append(fields, fmt.Sprintf("%s=%s", k, v))
	}

	fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("\x1b[%dm%s\x1b %s %s\t[%s] %s \n", red, strings.ToUpper(entry.Level.String()), timestamp, fileVal, entry.Message)
	//newLog = fmt.Sprintf("\x1b%d\x1b %s\t%s %s %s\t[%s] %s \n", red, strings.ToUpper(entry.Level.String()), timestamp, fileVal, entry.Message, strings.Join(fields, " "), getPackage())

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func getPackage() []byte {
	pc, _, _, _ := runtime.Caller(0)
	fullFuncName := runtime.FuncForPC(pc).Name()
	idx := strings.LastIndex(fullFuncName, ".")
	return []byte(fullFuncName[:idx]) // trim off function details
}

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

var (
	// qualified package name, cached at first use
	logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		//If the caller isn't part of this package, we're done
		if pkg == logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	//if !strings.Contains(f, "/") {
	//	return f
	//}
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
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
