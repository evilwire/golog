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
//	logger = golog.GetLogger("my-module")
//	logger.Infof("Hello: %s", someVar)
//
//	// the following will not do anything
//	logger.Debug("Goodbye!")
// }
//
// More elaborate example:
//
// golog.GetLogger("my-module").Info("Foobar!")
//
package golog

import "github.com/golang/glog"


// Configuration
//
type LogConfig struct {
	Prefix string
	Level level
}

// Logger
//
type Logger interface {
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Info(args ...interface{})

	Fatalf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Warningf(message string, args ...interface{})
	Infof(message string, args ...interface{})
}

// Abstract out the glog.Fatal and glog.Fatalf functions for
// unit testing
type logFun func (args ...interface{})
type logfFun func (message string, args ...interface{})


type logger struct {
	config LogConfig

	fatal logFun
	fatalf logfFun

	error logFun
	errorf logfFun

	warn logFun
	warnf logfFun

	info logFun
	infof logfFun
}


func (log *logger) log(l level, args ...interface{}) {
	if log.config.Level >= l {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.info(args...)
	}
}

func (log *logger) logf(l level, message string, args ...interface{}) {
	if log.config.Level >= l {
		message = log.config.Prefix + message
		log.infof(message, args...)
	}
}

func (log *logger) Fatal(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.fatal(args...)
}

func (log *logger) Fatalf(message string, args ...interface{}) {
	message = log.config.Prefix + message
	log.fatalf(message, args...)
}

func (log *logger) Error(args ...interface{}) {
	if log.config.Level >= ERROR {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.error(args...)
	}
}

func (log *logger) Errorf(message string, args ...interface{}) {
	if log.config.Level >= ERROR {
		message = log.config.Prefix + message
		log.errorf(message, args...)
	}
}

func (log *logger) Warn(args ...interface{}) {
	if log.config.Level >= WARN {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.warn(args...)
	}
}

func (log *logger) Warnf(message string, args ...interface{}) {
	if log.config.Level >= WARN {
		message = log.config.Prefix + message
		log.warnf(message, args...)
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

func (log *logger) Debugf(message string, args ...interface{}) {
	log.logf(DEBUG, message, args...)
}

func newLogger(config LogConfig) *logger {
	return &logger{
		config: config,
		fatal:  glog.Fatal,
		fatalf: glog.Fatalf,
		error:  glog.Error,
		errorf: glog.Errorf,
		warn:   glog.Warning,
		warnf:  glog.Warningf,
		info:    glog.Info,
		infof:   glog.Infof,
	}
}

var loggers map[string]*logger = make(map[string]*logger)

func GetLogger(name string) *logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	l := newLogger(LogConfig{
		Level: DEBUG,
		Prefix: "[" + name + "] ",
	})

	loggers[name] = l
	return l
}

func Setup(name string, logConfig LogConfig) {
	l := newLogger(logConfig)
	loggers[name] = l
}
