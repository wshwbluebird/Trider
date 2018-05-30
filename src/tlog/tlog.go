package tlog

import (
	"log"
	"os"
)

const STDOUT  = 0
const FILEOUT = 1

type Tlog struct{
	logType uint
	logger *log.Logger
	isopen bool
}


func NewFileLog(path string, name string) (*Tlog){
	filename := "log." + name
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic("no this path : " + path + "\n")
	}

	f, err := os.OpenFile(path+"/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("cannot open this log : " + path + "/" + filename + "\n")
	}

	tlog := &Tlog{}
	tlog.isopen = false
	tlog.logger = log.New(f,"",log.LstdFlags)
	tlog.logType = FILEOUT
	return tlog
}


func NewStdOut() (*Tlog){
	tlog := &Tlog{}
	tlog.isopen = false
	tlog.logger = log.New(os.Stdout,"",log.LstdFlags)
	tlog.logType = STDOUT
	return tlog
}


func (tlog *Tlog) log(prefix string, content string){
	tlog.logger.Printf("[%s]: %s\n",prefix,content)
}

// LogError logs error info.
func (tlog *Tlog) LogError(str string) {
	tlog.log("ERROR", str)
}

// LogError logs normal info.
func (tlog *Tlog) LogRun(str string) {
	tlog.log("RUN", str)
}


func (tlog *Tlog) LogComplete(str string) {
	tlog.log("COMPLETE", str)
}


func (tlog *Tlog) LogFail(str string) {
	tlog.log("FAIL", str)
}


func (tlog *Tlog) LogNewUrl(str string) {
	tlog.log("NEW", str)
}