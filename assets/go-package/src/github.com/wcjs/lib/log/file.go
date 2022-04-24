package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"fmt"
)

type Loggerconfig struct {
	DebugEnabled  bool
	Dir           string
	Loglevel      int64
	DistingType   int64

	/*
	FatalLevel 4
	ErrorLevel 3
	WarnLevel  2
	InfoLevel  1
	DebugLevel 0
	*/
}


func LogRun(logconfig *Loggerconfig) {
	Logconfig = logconfig
	initLogger()
}

var Logconfig *Loggerconfig
var LogAccess *logrus.Logger
var LogError *logrus.Logger

var LogAccessFile *os.File
var LogErrorFile *os.File

var stimenowname string = ""

func initLogger() {
	LogAccess = logrus.New()
	LogAccess.Formatter = &logrus.SeasFormatter{TimestampFormat:"2006-01-02 15:04:05"}
	LogError = logrus.New()
	LogError.Formatter =  &logrus.JSONFormatter{DisableTimestamp:false,TimestampFormat:"2006-01-02 15:04:05"}

	if Logconfig.DebugEnabled {
		LogAccess.Level = logrus.DebugLevel
		LogError.Level = logrus.DebugLevel
		LogAccess.Out = os.Stdout
		LogError.Out = os.Stdout
	} else {
		switch Logconfig.Loglevel {
		case 0:
			LogAccess.Level = logrus.DebugLevel
			LogError.Level = logrus.DebugLevel
		case 1:
			LogAccess.Level = logrus.InfoLevel
			LogError.Level = logrus.InfoLevel
		case 2:
			LogAccess.Level = logrus.WarnLevel
			LogError.Level = logrus.WarnLevel
		case 3:
			LogAccess.Level = logrus.ErrorLevel
			LogError.Level = logrus.ErrorLevel
		case 4:
			LogAccess.Level = logrus.DebugLevel
			LogError.Level = logrus.DebugLevel
		default:
			LogAccess.Level = logrus.DebugLevel
			LogError.Level = logrus.DebugLevel
		}
		openLogFile()
		go autoSwitchLogFile()
	}
}

func autoSwitchLogFile() {
	for {
		<-time.After(100 * time.Second)
		openLogFile()
	}
}

func openLogFile() {
	var stime string
	if Logconfig.DistingType == 1 { //hour
		stime = time.Now().Format("2006-01-02_15")
	} else { //day
		stime = time.Now().Format("2006-01-02")
	}
	if stimenowname != stime {
		var err error
		accesslogpath := Logconfig.Dir + "/access_" + stime + ".log"
		errorlogpath := Logconfig.Dir + "/error_" + stime + ".log"
		//LogAccess.Lock() //切换日志文件
		if stimenowname != "" {
			LogAccessFile.Close()
		}
		LogAccessFile, err = os.OpenFile(accesslogpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			LogAccess.Out = os.Stdout
			logrus.Fatal(err)
		} else {
			LogAccess.Out = LogAccessFile
		}
		//LogAccess.UnLock()

		//LogError.Lock() //切换日志文件
		if stimenowname != "" {
			LogErrorFile.Close()
		}
		LogErrorFile, err = os.OpenFile(errorlogpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			LogError.Out = os.Stdout
			logrus.Fatal(err)
		} else {
			LogError.Out = LogErrorFile
		}
		//LogError.UnLock()
		stimenowname = stime
	}
}

func FileError(s ...string) {
	l := logrus.Fields{}
	for k, v:= range s{
		l[fmt.Sprint(k)] = fmt.Sprint(v)
	}
	LogError.WithFields(l).Error(s)
}

func FileAccess(s ...string) {
	l := logrus.Fields{}
	for k, v:= range s{
		l[fmt.Sprint(k)] = fmt.Sprint(v)
	}
	LogAccess.WithFields(l).Info(s)
}

func WriteErrLog(msg string) {
	LogErrorFile.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + ":" + msg + "\n"))
}

func WriteAccLog(msg string) {
	LogAccessFile.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + ":" + msg + "\n"))
}
