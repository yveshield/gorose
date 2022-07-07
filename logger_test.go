package gorose

import (
	"testing"
	"time"
)

func TestDefaultLogger(t *testing.T) {
	l := NewLogger(&LogOption{FilePath: "/tmp/gorose.log", EnableErrorLog: true})
	var sqlstr = "select xxx from xxx where a='a' and b=\"33\""
	l.Sql(0, sqlstr, time.Duration(1<<4))
	t.Log("logger success")
}
