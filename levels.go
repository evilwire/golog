// Logging levels
//
//
package golog

import (
	"github.com/golang/glog"
	"strings"
)


type level glog.Level

const (
	NOLOG = level(-1)
	ERROR = level(1)
	WARN = level(2)
	VERBOSE = level(3)
	INFO = level(4)
	DEBUG = level(5)
)

func GetLevel(levelName string) (level, bool) {
	levelName = strings.ToUpper(levelName)

	switch (levelName) {
	case "NOLOG":
		return NOLOG, true

	case "ERROR":
		return ERROR, true

	case "WARN":
		return WARN, true

	case "VERBOSE":
		return VERBOSE, true

	case "INFO":
		return INFO, true

	case "DEBUG":
		return DEBUG, true
	}

	return NOLOG, false
}