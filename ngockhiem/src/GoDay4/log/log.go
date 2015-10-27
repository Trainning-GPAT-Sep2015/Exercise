package log

import (
	"bytes"
	"log"
	"os"
	"sync"
	"time"
)

const (
	DEBUGLEVEL    = "DEBUG"
	INFOLEVEL     = "INFO"
	WARNINGLEVEL  = "WARNING"
	ERRORLEVEL    = "ERROR"
	CRITICALLEVEL = "CRITICAL"
	FILENAME      = "log.txt"
	TIMEFORMAT    = "02-01-2006 15:04:05"
)

/**
Log is an interface

*/
type Log interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Critical(msg string)
}

// Stdout log info to Std
type StdOut struct {
	out    bytes.Buffer
	prefix string
}

func (this *StdOut) Debug(msg string) {
	this.StdLog(DEBUGLEVEL, msg)
}

func (this *StdOut) Info(msg string) {
	this.StdLog(INFOLEVEL, msg)
}

func (this *StdOut) Warning(msg string) {
	this.StdLog(WARNINGLEVEL, msg)
}

func (this *StdOut) Error(msg string) {
	this.StdLog(ERRORLEVEL, msg)
}

func (this *StdOut) Critical(msg string) {
	this.StdLog(CRITICALLEVEL, msg)
}

func (this *StdOut) StdLog(level, msg string) {
	_, err := this.out.WriteString(time.Now().Format(TIMEFORMAT) + " " + this.prefix + " [" + level + "] " + msg)
	if err != nil {
		panic(err)
	}
}

// FileLog log info to file
type FileLog struct {
	mutex  sync.Mutex
	out    bytes.Buffer
	prefix string
}

func (this *FileLog) Debug(msg string) {
	this.WriteLog(DEBUGLEVEL, msg)
}

func (this *FileLog) Info(msg string) {
	this.WriteLog(INFOLEVEL, msg)
}

func (this *FileLog) Warning(msg string) {
	this.WriteLog(WARNINGLEVEL, msg)
}

func (this *FileLog) Error(msg string) {
	this.WriteLog(ERRORLEVEL, msg)
}

func (this *FileLog) Critical(msg string) {
	this.WriteLog(CRITICALLEVEL, msg)
}

func (this *FileLog) WriteLog(level, msg string) {
	_, err := this.out.WriteString(time.Now().Format(TIMEFORMAT) + " " + this.prefix + " [" + level + "] " + msg + "\n")
	if err != nil {
		panic(err)
	}
}

// interface for log
type LogAdapter interface {
	Output()
}

type StdOutAdapter struct {
	*StdOut
}

func (this StdOutAdapter) Output() {
	log.Println(this.out.String())
}

type FileLogAdapter struct {
	*FileLog
}

func (this FileLogAdapter) Output() {
	f, err := os.OpenFile(FILENAME, os.O_RDWR|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Write(this.out.Bytes())
}

type Logger struct {
	adapters []LogAdapter
}

func (this *Logger) Init(adapters []LogAdapter) {
	for _, adapter := range adapters {
		this.adapters = append(this.adapters, adapter)
	}
}

func (this *Logger) Debug(msg string) {
	for _, adapter := range this.adapters {
		filelog, ok1 := adapter.(FileLogAdapter)
		stdlog, _ := adapter.(StdOutAdapter)
		if ok1 {
			filelog.Debug(msg)
		} else {
			stdlog.Debug(msg)
		}
		adapter.Output()
	}
}

func (this *Logger) Info(msg string) {
	for _, adapter := range this.adapters {
		filelog, ok1 := adapter.(FileLogAdapter)
		stdlog, _ := adapter.(StdOutAdapter)
		if ok1 {
			filelog.Info(msg)
		} else {
			stdlog.Info(msg)
		}
		adapter.Output()
	}
}

func (this *Logger) Warning(msg string) {
	for _, adapter := range this.adapters {
		filelog, ok1 := adapter.(FileLogAdapter)
		stdlog, _ := adapter.(StdOutAdapter)
		if ok1 {
			filelog.Warning(msg)
		} else {
			stdlog.Warning(msg)
		}
		adapter.Output()
	}
}

func (this *Logger) Error(msg string) {
	for _, adapter := range this.adapters {
		filelog, ok1 := adapter.(FileLogAdapter)
		stdlog, _ := adapter.(StdOutAdapter)
		if ok1 {
			filelog.Error(msg)
		} else {
			stdlog.Error(msg)
		}
		adapter.Output()
	}
}

func (this *Logger) Critical(msg string) {
	for _, adapter := range this.adapters {
		filelog, ok1 := adapter.(FileLogAdapter)
		stdlog, _ := adapter.(StdOutAdapter)
		if ok1 {
			filelog.Critical(msg)
		} else {
			stdlog.Critical(msg)
		}
		adapter.Output()
	}
}
