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


type LogConfig struct {
	Prefix string
	Level level
}

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

type Glogger struct {}

func (log *Glogger) Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func (log *Glogger) Fatalf(message string, args ...interface{}) {
	glog.Fatalf(message, args...)
}

func (log *Glogger) Error(args ...interface{}) {
	glog.Error(args...)
}

func (log *Glogger) Errorf(message string, args ...interface{}) {
	glog.Errorf(message, args...)
}

func (log *Glogger) Warning(args ...interface{}) {
	glog.Warning(args...)
}

func (log *Glogger) Warningf(message string, args ...interface{}) {
	glog.Warningf(message, args...)
}

func (log *Glogger) Info(args ...interface{}) {
	glog.Info(args...)
}

func (log *Glogger) Infof(message string, args ...interface{}) {
	glog.Infof(message, args...)
}

type logger struct {
	config LogConfig
	base Logger
}


func (log *logger) log(l level, args ...interface{}) {
	if log.config.Level < l {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Info(args...)
	}
}

func (log *logger) logf(l level, message string, args ...interface{}) {
	if log.config.Level < l {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Infof(message, args...)
	}
}

func (log *logger) Fatal(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.base.Fatal(args...)
}

func (log *logger) Fatalf(message string, args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.base.Fatalf(message, args...)
}

func (log *logger) Error(args ...interface{}) {
	if log.config.Level >= ERROR {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Error(args...)
	}
}

func (log *logger) Errorf(message string, args ...interface{}) {
	if log.config.Level >= ERROR {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Errorf(message, args...)
	}
}

func (log *logger) Warn(args ...interface{}) {
	if log.config.Level >= WARN {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Warning(args...)
	}
}

func (log *logger) Warnf(message string, args ...interface{}) {
	if log.config.Level >= WARN {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.base.Warningf(message, args...)
	}
}

func (log *logger) Verbose(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.log(VERBOSE, args...)
}

func (log *logger) Verbosef(message string, args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.logf(VERBOSE, message, args...)
}

func (log *logger) Info(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.log(INFO, args...)
}

func (log *logger) Infof(message string, args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.logf(INFO, message, args...)
}

func (log *logger) Debug(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.log(DEBUG, args...)
}

func newLogger(config LogConfig) *logger {
	return &logger{config: config}
}

var loggers map[string]*logger = make(map[string]*logger)

func GetLogger(name string) *logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	l := newLogger(LogConfig{
		Level: DEBUG,
		Prefix: "[" + name + "]",
	})

	loggers[name] = l
	return l
}

func Setup(name string, logConfig *LogConfig) {
	l := newLogger(*logConfig)
	loggers[name] = l
}
