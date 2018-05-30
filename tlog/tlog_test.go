package tlog

import "testing"

func TestNewStdOut(t *testing.T) {
	tlog := NewStdOut()
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogComplete("sds")
}


func TestNewFileLog(t *testing.T) {
	tlog := NewFileLog("logresultw","gg")
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogError("hello")
	tlog.LogComplete("sds")
}
