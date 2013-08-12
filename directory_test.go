package main

import (
	"fmt"
	"reflect"
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

func TestWalkDirectory(t *testing.T) {
	directoryTree := walkDirectory("testdir")

	assertEqual(t, 4, len(directoryTree))
	assertEqual(t, "testdir/subdir1/a.txt", directoryTree[0].Path)
	assertEqual(t, "testdir/subdir1/b.txt", directoryTree[1].Path)
	assertEqual(t, "testdir/subdir2/b.txt", directoryTree[2].Path)
	assertEqual(t, "testdir/subdir2/c.txt", directoryTree[3].Path)	
}