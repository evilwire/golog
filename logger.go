// Go wrapper for glog.
//
// What we want to extend:
//
// - being able to (concurrently) specify module level log config
// - being able to specify the prefix for the module
//
// Author: Knight W. Fu
// License: Licensed under all kinds of funny things
//
// Basic Example:
//
// func main() {
// 	golog.SetUp("my-module", &glog.LogConfig{
// 		Prefix: "[my-module]",
//		Level: golog.INFO,
//	}
//
// 	// more functions!
//	// ...
//	logger = golog.Logger("my-module")
//	logger.Infof("Hello: %s", someVar)
//
//	// the following will not do anything
//	logger.Debug("Goodbye!")
// }
//
// More elaborate example:
//
// golog.Logger("my-module").Info("Foobar!")
//
package golog

import "github.com/golang/glog"

type logger struct {
	level level
}

type LogConfig struct {
	Prefix string
	Level level
}

func (log *logger) log(l level, args ...interface{}) {
	if log.level < l {
		glog.Info(args...)
	}
}

func (log *logger) logf(l level, message string, args ...interface{}) {
	if log.level < l {
		glog.Infof(message, args...)
	}
}

func (log *logger) Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func (log *logger) Fatalf(message string, args ...interface{}) {
	glog.Fatalf(message, args...)
}

func (log *logger) Error(args ...interface{}) {
	if log.level >= ERROR {
		glog.Error(args...)
	}
}

func (log *logger) Errorf(message string, args ...interface{}) {
	if log.level >= ERROR {
		glog.Errorf(message, args...)
	}
}

func (log *logger) Warn(args ...interface{}) {
	if log.level >= WARN {
		glog.Warning(args...)
	}
}

func (log *logger) Warnf(message string, args ...interface{}) {
	if log.level >= WARN {
		glog.Warningf(message, args...)
	}
}

func (log *logger) Verbose(args ...interface{}) {
	log.log(VERBOSE, args...)
}

func (log *logger) Verbosef(message string, args ...interface{}) {
	log.logf(VERBOSE, message, args...)
}

func (log *logger) Info(args ...interface{}) {
	log.log(INFO, args...)
}

func (log *logger) Infof(message string, args ...interface{}) {
	log.logf(INFO, message, args...)
}

func (log *logger) Debug(args ...interface{}) {
	log.log(DEBUG, args...)
}

func newLogger(l level) *logger {
	return &logger{level: l}
}

var loggers map[string]*logger = make(map[string]*logger)


func Logger(name string) *logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	l := newLogger(DEBUG)
	loggers[name] = l
	return l
}


func Setup(name string, logConfig *LogConfig) {
	l := newLogger(logConfig.Level)
	loggers[name] = l
}
