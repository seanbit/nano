package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCaller(t *testing.T) {
	SetLogger(logrus.New())
	Rpc()
}

func Rpc()  {
	WithCaller().Error("haha")
}

