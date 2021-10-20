// Copyright (c) nano Authors. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

// Logger represents  the log interface
type Logger interface {
	//Println(v ...interface{})
	//Fatal(v ...interface{})
	//Fatalf(format string, v ...interface{})

	logrus.FieldLogger
}

func init() {
	SetLogger(logrus.New())
}

var (
	//Println func(v ...interface{})
	//Fatal   func(v ...interface{})
	//Fatalf  func(format string, v ...interface{})

	WithCaller func() Logger


	Debugf func(format string, args ...interface{})
	Infof func(format string, args ...interface{})
	Printf func(format string, args ...interface{})
	Errorf func(format string, args ...interface{})
	Fatalf func(format string, args ...interface{})

	Debug func(args ...interface{})
	Info func(args ...interface{})
	Print func(args ...interface{})
	Error func(args ...interface{})
	Fatal func(args ...interface{})

	Debugln func(args ...interface{})
	Infoln func(args ...interface{})
	Println func(args ...interface{})
	Errorln func(args ...interface{})
	Fatalln func(args ...interface{})
)

// SetLogger rewrites the default logger
func SetLogger(logger Logger) {
	if logger == nil {
		return
	}
	//Println = logger.Println
	//Fatal = logger.Fatal
	//Fatalf = logger.Fatalf

	WithCaller = func() Logger {
		return logger.WithField("caller", caller())
	}

	Debugf = logger.Debugf
	Infof = logger.Infof
	Printf = logger.Printf
	Errorf = logger.Errorf
	Fatalf = logger.Fatalf

	Debug = logger.Debug
	Info = logger.Info
	Print = logger.Print
	Error = logger.Error
	Fatal = logger.Fatal

	Debugln = logger.Debugln
	Infoln = logger.Infoln
	Println = logger.Println
	Errorln = logger.Errorln
	Fatalln = logger.Fatalln
}

func caller() string {
	pc, _, _, ok := runtime.Caller(3)
	if ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}