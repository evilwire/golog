package golog

import (
	"testing"
)


func TestGetLogger(t *testing.T) {
	// setup the loggers
	loggers = map[string]*logger {
		"A": newLogger(LogConfig {
			Prefix: "[A]",
			Level: NOLOG,
		}),

		"B": newLogger(LogConfig {
			Prefix: "[B Logger]",
			Level: INFO,
		}),
	}

	// retrieve the test cases
	testCases := []struct {
		LoggerName string
		Level level
		Prefix string
	}{
		{
			LoggerName: "A",
			Level: NOLOG,
			Prefix: "[A]",
		},
		{
			LoggerName: "B",
			Level: INFO,
			Prefix: "[B Logger]",
		},
		{
			LoggerName: "C",
			Level: DEBUG,
			Prefix: "[C] ",
		},
	}

	for i, c := range testCases {
		log := GetLogger(c.LoggerName)
		if log.config.Level != c.Level {
			t.Errorf("TC %d: Expected level to be %d, actual: %d",
				i,
				c.Level,
				log.config.Level,
			)
		}

		if log.config.Prefix != c.Prefix {
			t.Errorf("TC %d: Expected prefix to be %s, actual: %s",
				i,
				c.Prefix,
				log.config.Prefix,
			)
		}
	}
}

func TestSetup(t *testing.T) {
	testCases := []struct {
		Name string
		Config LogConfig
	} {
		{

		},
	}

	for _, c := range testCases {
		Setup(c.Name, c.Config)
	}

	for i, c := range testCases {
		if _, ok := loggers[c.Name]; !ok {
			t.Errorf("TC %d: expected logger with name %s to exist",
				c.Name,
				i,
			)
		}

		log := GetLogger(c.Name)

		if log.config.Prefix != c.Config.Prefix {
			t.Errorf("TC %d: expected logger to have prefix %s, got %s instead",
				c.Config.Prefix,
				log.config.Prefix,
				i,
			)
		}

		if log.config.Level != c.Config.Level {
			t.Errorf("TC %d: expected logger to have level %d, got %d instead",
				c.Config.Level,
				log.config.Level,
				i,
			)
		}
	}
}
