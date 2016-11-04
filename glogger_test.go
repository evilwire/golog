package golog

import (
	"testing"
	"github.com/golang/glog"
	"reflect"
)


func isSameArrays(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}


type MockLogger struct {
	Message string
	Args []interface{}
	CallType string
}

func (logger *MockLogger) Fatal(args ...interface{}) {
	logger.CallType = "Fatal"
	logger.Args = args
}

func (logger *MockLogger) Fatalf(message string, args ...interface{}) {
	logger.CallType = "Fatalf"
	logger.Message = message
	logger.Args = args
}

func (logger *MockLogger) Error(args ...interface{}) {
	logger.CallType = "Error"
	logger.Args = args
}

func (logger *MockLogger) Errorf(message string, args ...interface{}) {
	logger.CallType = "Errorf"
	logger.Message = message
	logger.Args = args
}

func (logger *MockLogger) Warn(args ...interface{}) {
	logger.CallType = "Warn"
	logger.Args = args
}

func (logger *MockLogger) Warnf(message string, args ...interface{}) {
	logger.CallType = "Warnf"
	logger.Message = message
	logger.Args = args
}

func (logger *MockLogger) Info(args ...interface{}) {
	logger.CallType = "Info"
	logger.Args = args
}

func (logger *MockLogger) Infof(message string, args ...interface{}) {
	logger.CallType = "Infof"
	logger.Message = message
	logger.Args = args
}

func newLoggerWithMocks(config LogConfig) (*logger, *MockLogger) {
	mockLogger := MockLogger{
		Args: make([]interface{}, 0, 0),
	}

	return &logger {
		config: config,
		fatal: mockLogger.Fatal,
		fatalf: mockLogger.Fatalf,
		error: mockLogger.Error,
		errorf: mockLogger.Errorf,
		warn: mockLogger.Warn,
		warnf: mockLogger.Warnf,
		info: mockLogger.Info,
		infof: mockLogger.Infof,
	}, &mockLogger
}


func TestLogger_Fatal(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: NOLOG,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Fatal(c.Messages...)

		if mockLogger.CallType != "Fatal" {
			t.Errorf("TC %d: Expect the call type to be \"Fatal\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}

func TestLogger_Fatalf(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: NOLOG,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Fatalf(c.Template, c.Messages...)

		if mockLogger.CallType != "Fatalf" {
			t.Errorf("Test case %d: Expect the call type to be \"Fatalf\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipError(t *testing.T) {
	logConfig := LogConfig {
		Level: NOLOG,
		Prefix: "[TestLogger]",
	}

	log, mockLogger := newLoggerWithMocks(logConfig)
	log.Error("hello", "goodbye")

	if mockLogger.CallType != "" {
		t.Error("Expect Error to not be called.")
	}

	if len(mockLogger.Args) != 0 {
		t.Error("Expect Error to not be called.")
	}
}


func TestLogger_Error(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: ERROR,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Error(c.Messages...)

		if mockLogger.CallType != "Error" {
			t.Errorf("TC %d: Expect the call type to be \"Error\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}

func TestLogger_SkipErrorf(t *testing.T) {
	logConfig := LogConfig {
		Level: NOLOG,
		Prefix: "[TestLogger]",
	}

	log, mockLogger := newLoggerWithMocks(logConfig)
	log.Errorf("%s: %s", "hello", "goodbye")

	if mockLogger.CallType != "" {
		t.Error("Expect Error to not be called.")
	}

	if len(mockLogger.Args) != 0 {
		t.Error("Expect Error to not be called.")
	}
}

func TestLogger_Errorf(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Errorf(c.Template, c.Messages...)

		if mockLogger.CallType != "Errorf" {
			t.Errorf("Test case %d: Expect the call type to be \"Errorf\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipWarn(t *testing.T) {

	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipWarn]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipWarn]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Warn("hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Warn not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Warn not to be called.", i)
		}
	}
}


func TestLogger_Warn(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: WARN,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Warn(c.Messages...)

		if mockLogger.CallType != "Warn" {
			t.Errorf("TC %d: Expect the call type to be \"Warn\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}

func TestLogger_SkipWarnf(t *testing.T) {
	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipWarnf]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipWarnf]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Warnf("%s: %s", "hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Warnf not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Warnf not to be called.", i)
		}
	}
}

func TestLogger_Warnf(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger_Warnf]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger_Warnf]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Warnf(c.Template, c.Messages...)

		if mockLogger.CallType != "Warnf" {
			t.Errorf("Test case %d: Expect the call type to be \"Warnf\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipVerbose(t *testing.T) {

	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: DEBUG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Verbose("hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}


func TestLogger_Verbose(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Verbose(c.Messages...)

		if mockLogger.CallType != "Info" {
			t.Errorf("TC %d: Expect the call type to be \"Info\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}

func TestLogger_SkipVerbosef(t *testing.T) {
	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: DEBUG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Verbosef("%s: %s", "hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}

func TestLogger_Verbosef(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: VERBOSE,
				Prefix: "[TestLogger-DEBUG]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Verbosef(c.Template, c.Messages...)

		if mockLogger.CallType != "Infof" {
			t.Errorf("Test case %d: Expect the call type to be \"Infof\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipInfo(t *testing.T) {

	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: WARN,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Info("hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}


func TestLogger_Info(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Info(c.Messages...)

		if mockLogger.CallType != "Info" {
			t.Errorf("TC %d: Expect the call type to be \"Info\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}

func TestLogger_SkipInfof(t *testing.T) {
	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: WARN,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Infof("%s: %s", "hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}

func TestLogger_Infof(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: INFO,
				Prefix: "[TestLogger-DEBUG]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Infof(c.Template, c.Messages...)

		if mockLogger.CallType != "Infof" {
			t.Errorf("Test case %d: Expect the call type to be \"Infof\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipDebug(t *testing.T) {

	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: INFO,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Debug("hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}


func TestLogger_Debug(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
	}{
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "",
			},
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "Non-CIISA Préfix",
			},
			Messages: []interface{}{},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Debug(c.Messages...)

		if mockLogger.CallType != "Info" {
			t.Errorf("TC %d: Expect the call type to be \"Info\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		if actualPrefix, ok := mockLogger.Args[0].(string); !ok {
			t.Errorf("TC %d: The first argument passed to the method should be string!",
				i,
			)
		} else if actualPrefix != c.Config.Prefix {
			t.Errorf("TC %d: Expects prefix %s, received %s instead",
				i,
				c.Config.Prefix,
				actualPrefix,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args[1:]) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestLogger_SkipDebugf(t *testing.T) {
	testCases := []LogConfig {
		{
			Level: NOLOG,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: ERROR,
			Prefix: "[TestLogger_SkipVerbose]",
		},
		{
			Level: INFO,
			Prefix: "[TestLogger_SkipVerbose]",
		},
	}


	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c)
		log.Debugf("%s: %s", "hello", "goodbye")

		if mockLogger.CallType != "" {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}

		if len(mockLogger.Args) != 0 {
			t.Errorf("TC %d: Expect Verbose not to be called.", i)
		}
	}
}


func TestLogger_Debugf(t *testing.T) {
	testCases := []struct {
		Config LogConfig
		Messages []interface{}
		Template string
	}{
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger]",
			},
			Messages: []interface{}{
				"Hello!",
			},
			Template: "%s",
		},
		{
			Config: LogConfig {
				Level: DEBUG,
				Prefix: "[TestLogger-DEBUG]",
			},
			Template: "%s",
			Messages: []interface{}{
				"Hello!", 1, true, struct{ A string }{"a"},
			},
		},
	}

	for i, c := range testCases {
		log, mockLogger := newLoggerWithMocks(c.Config)
		log.Debugf(c.Template, c.Messages...)

		if mockLogger.CallType != "Infof" {
			t.Errorf("Test case %d: Expect the call type to be \"Infof\", got \"%s\" instead",
				i,
				mockLogger.CallType,
			)
		}

		expectedTemplate := c.Config.Prefix + c.Template
		if expectedTemplate != mockLogger.Message {
			t.Errorf("TC %d: The templating message are not the same! Expect %s, Actual %s",
				i,
				expectedTemplate,
				mockLogger.Message,
			)
		}

		if !isSameArrays(c.Messages, mockLogger.Args) {
			t.Errorf("TC %d: The inputs are not the same!",
				i,
			)
		}
	}
}


func TestNewLogger(t *testing.T) {
	logger := newLogger(LogConfig {
		Prefix: "[hello]",
		Level: INFO,
	})

	// check the log config
	if logger.config.Prefix != "[hello]" {
		t.Errorf("Expect config to be the same! Expect prefix %s, got %s instead",
			"[hello]",
			logger.config.Prefix,
		)
	}

	if logger.config.Level != INFO {
		t.Errorf("Expect config to be the same! Expect log level %d, got %d instead",
			INFO,
			logger.config.Level,
		)
	}

	// check the function pointers
	loggerFatal := reflect.ValueOf(logger.fatal)
	glogFatal := reflect.ValueOf(glog.Fatal)

	if (glogFatal.Pointer() != loggerFatal.Pointer()) {
		t.Error("Expects the fatal method to be the same as glog.Fatal.")
	}

	loggerFatalf := reflect.ValueOf(logger.fatalf)
	glogFatalf := reflect.ValueOf(glog.Fatalf)

	if (glogFatalf.Pointer() != loggerFatalf.Pointer()) {
		t.Error("Expects the fatalf method to be the same as glog.Fatalf.")
	}

	loggerError := reflect.ValueOf(logger.error)
	glogError := reflect.ValueOf(glog.Error)

	if (glogError.Pointer() != loggerError.Pointer()) {
		t.Error("Expects the error method to be the same as glog.Error!")
	}

	loggerErrorf := reflect.ValueOf(logger.errorf)
	glogErrorf := reflect.ValueOf(glog.Errorf)

	if (glogErrorf.Pointer() != loggerErrorf.Pointer()) {
		t.Error("Expects the errorf method to be the same as glog.Errorf!")
	}

	loggerWarn := reflect.ValueOf(logger.warn)
	glogWarn := reflect.ValueOf(glog.Warning)

	if (glogWarn.Pointer() != loggerWarn.Pointer()) {
		t.Error("Expects the warn method to be the same as glog.Warning!")
	}

	loggerWarnf := reflect.ValueOf(logger.warnf)
	glogWarnf := reflect.ValueOf(glog.Warningf)

	if (glogWarnf.Pointer() != loggerWarnf.Pointer()) {
		t.Error("Expects the warnf method to be the same as glog.Warningf!")
	}

	loggerInfo := reflect.ValueOf(logger.info)
	glogInfo := reflect.ValueOf(glog.Info)

	if (glogInfo.Pointer() != loggerInfo.Pointer()) {
		t.Error("Expects the info method to be the same as glog.Info!")
	}

	loggerInfof := reflect.ValueOf(logger.infof)
	glogInfof := reflect.ValueOf(glog.Infof)

	if (glogInfof.Pointer() != loggerInfof.Pointer()) {
		t.Error("Expects the infof method to be the same as glog.Infof!")
	}
}
