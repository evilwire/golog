package golog

import "testing"

func TestGetLevel(t *testing.T) {
	testCases := []struct {
		Input string
		Exists bool
		Level level
	}{
		{
			"NOLOG",
			true,
			NOLOG,
		},
		{
			"ERROR",
			true,
			ERROR,
		},
		{
			"WARN",
			true,
			WARN,
		},
		{
			"VERBOSE",
			true,
			VERBOSE,
		},
		{
			"INFO",
			true,
			INFO,
		},
		{
			"DEBUG",
			true,
			DEBUG,
		},
		{
			"debug",
			true,
			DEBUG,
		},
		{
			"UNKNOWN",
			false,
			NOLOG,
		},
	}

	for _, c := range testCases {
		l, exists := GetLevel(c.Input)
		if exists != c.Exists {
			errMessage := "Expect level to exist"
			if !c.Exists {
				errMessage = "Expect level to not exist"
			}

			t.Error(errMessage)
		}

		if l != c.Level {
			t.Errorf("Expected level %d, actual level %d",
				c.Level,
				l,
			)
		}
	}
}
