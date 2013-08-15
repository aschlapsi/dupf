package dupfinder

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	fi := NewFileInfo("testdir/subdir1/a.txt")

	assertSuffix(t, "testdir/subdir1/a.txt", fi.path)
	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fi.GetHashstring())
}

func TestWalkDirectory(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, 5, dirTree.FileCount())
	assertSuffix(t, "testdir/subdir1/a.txt", dirTree.files[0].path)
	assertSuffix(t, "testdir/subdir1/b.txt", dirTree.files[1].path)
	assertSuffix(t, "testdir/subdir2/b.txt", dirTree.files[2].path)
	assertSuffix(t, "testdir/subdir2/c.txt", dirTree.files[3].path)
	assertSuffix(t, "testdir/subdir2/d.txt", dirTree.files[4].path)
}

func TestFilePathsAreAbsolute(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertTrue(t, filepath.IsAbs(dirTree.files[0].path), "Filepath must be absolute.")
	assertTrue(t, filepath.IsAbs(dirTree.files[1].path), "Filepath must be absolute.")
	assertTrue(t, filepath.IsAbs(dirTree.files[2].path), "Filepath must be absolute.")
	assertTrue(t, filepath.IsAbs(dirTree.files[3].path), "Filepath must be absolute.")
	assertTrue(t, filepath.IsAbs(dirTree.files[4].path), "Filepath must be absolute.")
}

func TestHashCodes(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", fmt.Sprintf("%x", dirTree.files[0].hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.files[1].hash))
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", fmt.Sprintf("%x", dirTree.files[2].hash))
	assertEqual(t, "4a8a08f09d37b73795649038408b5f33", fmt.Sprintf("%x", dirTree.files[3].hash))
	assertEqual(t, "d4c7ede6154c1efe72fd8b10cac048b0", fmt.Sprintf("%x", dirTree.files[4].hash))
}

func TestGetHashstring(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", dirTree.files[0].GetHashstring())
}

func TestFileSizes(t *testing.T) {
	dirTree := NewDirTree("testdir")

	assertEqual(t,  1, dirTree.files[0].size)
	assertEqual(t,  1, dirTree.files[1].size)
	assertEqual(t,  1, dirTree.files[2].size)
	assertEqual(t,  1, dirTree.files[3].size)
	assertEqual(t, 12, dirTree.files[4].size)
}

func TestAppendDir(t *testing.T) {
	dirTree := NewDirTree("testdir/subdir1")

	dirTree.AppendDir("testdir/subdir2")

	assertEqual(t, 5, dirTree.FileCount())
	assertSuffix(t, "testdir/subdir1/a.txt", dirTree.files[0].path)
	assertSuffix(t, "testdir/subdir1/b.txt", dirTree.files[1].path)
	assertSuffix(t, "testdir/subdir2/b.txt", dirTree.files[2].path)
	assertSuffix(t, "testdir/subdir2/c.txt", dirTree.files[3].path)
	assertSuffix(t, "testdir/subdir2/d.txt", dirTree.files[4].path)
}