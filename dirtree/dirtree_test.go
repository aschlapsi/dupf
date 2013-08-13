package dirtree

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

func TestNewFileInfo(t *testing.T) {
	fi := NewFileInfo("testdir/subdir1/a.txt")

	assertEqual(t, "testdir/subdir1/a.txt", fi.Path)
	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fi.GetHashstring())
}

func TestWalkDirectory(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, 4, dirTree.FileCount())
	assertEqual(t, "testdir/subdir1/a.txt", dirTree.Files[0].Path)
	assertEqual(t, "testdir/subdir1/b.txt", dirTree.Files[1].Path)
	assertEqual(t, "testdir/subdir2/b.txt", dirTree.Files[2].Path)
	assertEqual(t, "testdir/subdir2/c.txt", dirTree.Files[3].Path)
}

func TestHashCodes(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fmt.Sprintf("%x", dirTree.Files[0].Hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.Files[1].Hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.Files[2].Hash))
	assertEqual(t, "4a8a08f09d37b73795649038408b5f33", fmt.Sprintf("%x", dirTree.Files[3].Hash))
}

func TestGetHashstring(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", dirTree.Files[0].GetHashstring())
}
