// Copyright 2014 Christoffer Hallas.

package stacko

import (
	"testing"
)

// TestNewStacktrace creates a new stacktrace skipping 0 frames. If the creation
// of the stacktrace is correct, this should gives us a stacktrace of 4 frames.
func TestNewStacktrace(t *testing.T) {
	stacktrace, err := NewStacktrace(0)
	var expected string
	if err != nil {
		t.Error(err)
	}

	if len(stacktrace) != 4 {
		t.Error("Stacktrace should be 4 frames.")
	}

	expected = "stacko"
	if stacktrace[0].PackageName != expected {
		t.Errorf("Expected stacktrace[0].PackageName == '%s'. Was: '%s'",
			expected, stacktrace[0].PackageName)
	}

	expected = "NewStacktrace"
	if stacktrace[0].FunctionName != expected {
		t.Errorf("Expected stacktrace[0].FunctionName == '%s'. Was: '%s'",
			expected, stacktrace[0].FunctionName)
	}

	expected = "stacko"
	if stacktrace[1].PackageName != expected {
		t.Errorf("Expected stacktrace[1].PackageName == '%s'. Was: '%s'",
			expected, stacktrace[1].PackageName)
	}

	expected = "TestNewStacktrace"
	if stacktrace[1].FunctionName != expected {
		t.Errorf("Expected stacktrace[1].FunctionName == '%s'. Was: '%s'",
			expected, stacktrace[1].FunctionName)
	}

	expected = "testing"
	if stacktrace[2].PackageName != expected {
		t.Errorf("Expected stacktrace[2].PackageName == '%s'. Was: '%s'",
			expected, stacktrace[2].PackageName)
	}

	expected = "tRunner"
	if stacktrace[2].FunctionName != expected {
		t.Errorf("Expected stacktrace[2].FunctionName == '%s'. Was: '%s'",
			expected, stacktrace[2].FunctionName)
	}

	expected = "runtime"
	if stacktrace[3].PackageName != expected {
		t.Errorf("Expected stacktrace[3].PackageName == '%s'. Was: '%s'",
			expected, stacktrace[3].PackageName)
	}

	expected = "goexit"
	if stacktrace[3].FunctionName != expected {
		t.Errorf("Expected stacktrace[3].FunctionName == '%s'. Was: '%s'",
			expected, stacktrace[3].FunctionName)
	}
}
