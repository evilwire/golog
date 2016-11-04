// Package representing Logging levels (enum type)
package golog

import (
	"github.com/golang/glog"
	"strings"
)

// An alias for glog.Level, declares the enum type but restricts the user
// from being able to create instances
type level glog.Level

// Enum representing log levels
const (

	// do not log except fatal
	NOLOG = level(-1)

	// logging only errors
	ERROR = level(1)

	// logging errors and warning
	WARN = level(2)

	// first level of text-based logging
	VERBOSE = level(3)

	// second level of text-based logging
	INFO = level(4)

	// third and last level of text-based logging
	DEBUG = level(5)
)

// Retrieve a certain level by name, and if the name is not recognised returns NOLOG,
// false to indicate that the level is not recognized.
//
// levelName (string) - the name of the level
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