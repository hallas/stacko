package stacko

import (
	"testing"
)

// TestNewStacktrace creates a new stacktrace skipping 0 frames. If the creation
// of the stacktrace is correct, this should gives us a stacktrace of 4 frames.
func TestNewStacktrace(t *testing.T) {
	stacktrace, err := NewStacktrace(0)
	if err != nil {
		t.Error(err)
	}

	if len(stacktrace) != 4 {
		t.Error("Stacktrace should be 4 frames.")
	}

	var testFrame = func(frame int, v, expected string) {
		if v != expected {
			t.Errorf("Expected stacktrace[%d].PackageName == '%s'. Was: '%s'",
				frame, expected, v)
		}
	}

	testFrame(0, stacktrace[0].PackageName, "stacko")
	testFrame(0, stacktrace[0].FunctionName, "NewStacktrace")
	testFrame(1, stacktrace[1].PackageName, "stacko")
	testFrame(1, stacktrace[1].FunctionName, "TestNewStacktrace")
	testFrame(2, stacktrace[2].PackageName, "testing")
	testFrame(2, stacktrace[2].FunctionName, "tRunner")
	testFrame(3, stacktrace[3].PackageName, "runtime")
	testFrame(3, stacktrace[3].FunctionName, "goexit")
}
