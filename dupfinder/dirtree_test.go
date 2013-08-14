package dupfinder

import (
	"fmt"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	fi := NewFileInfo("testdir/subdir1/a.txt")

	assertEqual(t, "testdir/subdir1/a.txt", fi.path)
	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fi.GetHashstring())
}

func TestWalkDirectory(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, 4, dirTree.FileCount())
	assertEqual(t, "testdir/subdir1/a.txt", dirTree.files[0].path)
	assertEqual(t, "testdir/subdir1/b.txt", dirTree.files[1].path)
	assertEqual(t, "testdir/subdir2/b.txt", dirTree.files[2].path)
	assertEqual(t, "testdir/subdir2/c.txt", dirTree.files[3].path)
}

func TestHashCodes(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fmt.Sprintf("%x", dirTree.files[0].hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.files[1].hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.files[2].hash))
	assertEqual(t, "4a8a08f09d37b73795649038408b5f33", fmt.Sprintf("%x", dirTree.files[3].hash))
}

func TestGetHashstring(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", dirTree.files[0].GetHashstring())
}
