package main

import (
	"./log"
	"bytes"
	"fmt"
	"os"
	"time"
)

var (
	a log.Log
)

func main() {
	var a bytes.Buffer
	a.WriteString(time.Now().Format("02-01-2006 15:04:05") + " prefix" + " [" + "level" + "] " + "msg\n")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Write(a.Bytes())
	var std log.Logger
	std.Init([]log.LogAdapter{log.FileLogAdapter})
	fmt.Println(std)
}
