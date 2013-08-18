package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	search := FindDuplicates("dupfinder/testdir")

	assertSuffix(t, "dupfinder/testdir/subdir1/a.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir1/b.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/b.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/c.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/d.txt", <-search.Progress())
	assertSuffix(t, "", <-search.Progress())
}

func TestFilePathsAreAbsolute(t *testing.T) {
	search := FindDuplicates("dupfinder/testdir")

	assertTrue(t, filepath.IsAbs(<-search.Progress()), "File path is not absolute.")
	assertTrue(t, filepath.IsAbs(<-search.Progress()), "File path is not absolute.")
	assertTrue(t, filepath.IsAbs(<-search.Progress()), "File path is not absolute.")
	assertTrue(t, filepath.IsAbs(<-search.Progress()), "File path is not absolute.")
	assertTrue(t, filepath.IsAbs(<-search.Progress()), "File path is not absolute.")
}

func TestWalksAllDirectories(t *testing.T) {
	search := FindDuplicates("dupfinder/testdir/subdir1", "dupfinder/testdir/subdir2")

	assertSuffix(t, "dupfinder/testdir/subdir1/a.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir1/b.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/b.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/c.txt", <-search.Progress())
	assertSuffix(t, "dupfinder/testdir/subdir2/d.txt", <-search.Progress())
	assertSuffix(t, "", <-search.Progress())
}

func TestFiles(t *testing.T) {
	search := findTestDuplicates()

	assertEqual(t, 5, len(search.files))
	assertSuffix(t, "dupfinder/testdir/subdir1/a.txt", search.files[0].path)
	assertSuffix(t, "dupfinder/testdir/subdir1/b.txt", search.files[1].path)
	assertSuffix(t, "dupfinder/testdir/subdir2/b.txt", search.files[2].path)
	assertSuffix(t, "dupfinder/testdir/subdir2/c.txt", search.files[3].path)
	assertSuffix(t, "dupfinder/testdir/subdir2/d.txt", search.files[4].path)
}

func TestHashcodes(t *testing.T) {
	search := findTestDuplicates()

	assertEqual(t, 5, len(search.files))
	assertEqual(t, "0cc175b9c0f1b6a831c399e269772661", search.files[0].Hashstring())
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", search.files[1].Hashstring())
	assertEqual(t, "92eb5ffee6ae2fec3ad71c777531578f", search.files[2].Hashstring())
	assertEqual(t, "4a8a08f09d37b73795649038408b5f33", search.files[3].Hashstring())
	assertEqual(t, "d4c7ede6154c1efe72fd8b10cac048b0", search.files[4].Hashstring())
}

func TestTotalFileSize(t *testing.T) {
	search := findTestDuplicates()

	assertEqual(t, 1 + 1 + 1 + 1 + 12, search.TotalFileSize())
}

func TestTotalFileCount(t *testing.T) {
	search := findTestDuplicates()

	assertEqual(t, 5, search.FileCount())
}

func TestFoundDuplicates(t *testing.T) {
	search := findTestDuplicates()

	assertEqual(t, 1, search.Duplicates().Count())
	assertEqual(t, 2, search.Duplicates().TotalFileSize())
	assertSuffix(t, "dupfinder/testdir/subdir1/b.txt", search.Duplicates().Groups()[0][0].Path())
	assertSuffix(t, "dupfinder/testdir/subdir2/b.txt", search.Duplicates().Groups()[0][1].Path())
}

func findTestDuplicates() *SearchProgress {
	search := FindDuplicates("dupfinder/testdir")

	for {
		if (<-search.Progress()) == "" {
			break
		}
	}

	return search
}

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
