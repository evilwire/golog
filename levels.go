// Logging levels
//
//
package golog

import "github.com/golang/glog"


type level glog.Level

const (
	ERROR = level(1)
	WARN = level(2)
	VERBOSE = level(3)
	INFO = level(4)
	DEBUG = level(5)
)