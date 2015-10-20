package mylog

import (
	"time"
)

type 

type mylog struct {
}

type logAdapters []mylog

type LogAdapter interface {
	Init()
	New(moduleName string) LogAdapter
	GetLast(n) []LogAdapter
}

type StdOutAdapter struct {
	moduleName string
	outputs    string
}

type FileAdapter struct {
	moduleName string
}

func (s StdOutAdapter) Init() StdOutAdapter {
	if s == nil {
		s = new(StdOutAdapter)
	}
	return s
}

func (s StdOutAdapter) New() {
}

func (f FileAdapter) Init() FileAdapter {
	if f == nil {
		f = new(FileAdapter)
	}
	return f
}

func (f FileAdapter) New() {

}
