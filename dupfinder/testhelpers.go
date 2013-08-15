package dupfinder

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func areObjectsEqual(a, b interface{}) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	if reflect.ValueOf(a) == reflect.ValueOf(b) {
		return true
	}

	if fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b) {
		return true
	}

	return false
}

func assertEqual(t *testing.T, a, b interface{}) {
	if !areObjectsEqual(a, b) {
		t.Errorf("Not equal (expected: '%#v', but was: '%#v')", a, b)
	}
}

func assertSuffix(t *testing.T, a, b string) {
	if !strings.HasSuffix(b, a) {
		t.Errorf("Expected string to end with %s (string: %s)", a, b)
	}
}

func assertTrue(t *testing.T, r bool, message string) {
	if !r {
		t.Error(message)
	}
}