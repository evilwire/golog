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
// 	golog.Setup("my-module", &glog.LogConfig{
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
package golog

import "github.com/golang/glog"


// Represents the log configuration, and contains a way to configure the
// prefix and log level.
type LogConfig struct {
	// string that would occur at the beginning of every log
	Prefix string

	// the level of the log
	Level level
}


// Abstract out the log and logf functions for unit testing
// Represents the log function
type logFun func (args ...interface{})

// Represents the log + format function
type logfFun func (message string, args ...interface{})


// The base logger class; cannot be instantiated except through
// GetLogger.
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

// Logs arguments with the "fatal" declaration and exists with code 255.
// Logs will have an "F" at the beginning, and include the line no. where
// the log is issued.
func (log *logger) Fatal(args ...interface{}) {
	args = append([]interface{}{log.config.Prefix}, args...)
	log.fatal(args...)
}

// Logs a templated message with the "fatal" declaration and exists with
// code 255. Same as Fatal, except the arguments will be used to instantiate
// a templating string.
func (log *logger) Fatalf(message string, args ...interface{}) {
	message = log.config.Prefix + message
	log.fatalf(message, args...)
}

// Logs arguments with the "error" declaration. Logs will have an "E" at the
// beginning, and include the line no.
func (log *logger) Error(args ...interface{}) {
	if log.config.Level >= ERROR {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.error(args...)
	}
}

// Logs a templated message with the "error" declaration, same as Error,
// except formats the argument according to the message template.
func (log *logger) Errorf(message string, args ...interface{}) {
	if log.config.Level >= ERROR {
		message = log.config.Prefix + message
		log.errorf(message, args...)
	}
}

// Logs arguments with the "warning" declaration. Begins with a "W" and
// includes the line no.
func (log *logger) Warn(args ...interface{}) {
	if log.config.Level >= WARN {
		args = append([]interface{}{log.config.Prefix}, args...)
		log.warn(args...)
	}
}

// Logs a templated message with the "warning" declaration. Like Warn,
// except formats the args according to templates.
func (log *logger) Warnf(message string, args ...interface{}) {
	if log.config.Level >= WARN {
		message = log.config.Prefix + message
		log.warnf(message, args...)
	}
}

// Logs arguments with the "info" designation. Starts with an "I". This should
// be used widely for delineating high-level checkpoints for code paths. Any
// production level software should run with this level, and should not be used in
// noisy log dumps or for expensive computations.
func (log *logger) Info(args ...interface{}) {
	log.log(INFO, args...)
}

// Logs a templated message with the "info" designation.
func (log *logger) Infof(message string, args ...interface{}) {
	log.logf(INFO, message, args...)
}

// Logs arguments with the "debug" designation. Log messages start with an "I".
// This should be used for detailed information that can be used for debugging
// somewhat expensive code paths.
func (log *logger) Debug(args ...interface{}) {
	log.log(DEBUG, args...)
}

// Logs a templated message with the "debug" designation.
func (log *logger) Debugf(message string, args ...interface{}) {
	log.logf(DEBUG, message, args...)
}

// Logs messages as "verbose". Should be used for noisy logs, e.g. processing-level
// logs at the lowest of possible levels of the code.
func (log *logger) Verbose(args ...interface{}) {
	log.log(VERBOSE, args...)
}

// Logs a templated message with the "verbose" designation.
func (log *logger) Verbosef(message string, args ...interface{}) {
	log.logf(VERBOSE, message, args...)
}

//
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

// Get a logger by name. If the logger has not been previously setup
// the logger will be configured (and setup) with default level of "DEBUG"
// and the default prefix of "[$name] "
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

// Sets up a logger by name, and with a set of log configurations. This
// should be called at the start of the application.
func Setup(name string, logConfig LogConfig) {
	l := newLogger(logConfig)
	loggers[name] = l
}
