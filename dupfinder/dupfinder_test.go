package dupfinder

import (
	"testing"
)

func TestNewDuplicates(t *testing.T) {
	files := []*FileInfo{
		{path: "root/subdir1/a.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir1/b.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}, size: 1},
		{path: "root/subdir2/c.txt", hash: []byte{0x3, 0x4, 0x5, 0x6}, size: 1},
		{path: "root/subdir2/d.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}, size: 1},
		{path: "root/subdir3/e.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir3/f.txt", hash: []byte{0x4, 0x5, 0x6, 0x7}, size: 1},
		{path: "root/subdir4/g.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir4/h.txt", hash: []byte{0x5, 0x6, 0x7, 0x8}, size: 1},
	}

	dirTree := DirTree{rootPath: "root", files: files}

	duplicates := NewDuplicates(&dirTree)

	assertEqual(t, 2, duplicates.Count())
	assertEqual(t, "010203041", duplicates.items[0].key)
	assertEqual(t, 3, len(duplicates.items[0].files))
	assertEqual(t, files[0], duplicates.items[0].files[0])
	assertEqual(t, files[4], duplicates.items[0].files[1])
	assertEqual(t, files[6], duplicates.items[0].files[2])
	assertEqual(t, "020304051", duplicates.items[1].key)
	assertEqual(t, 2, len(duplicates.items[1].files))
	assertEqual(t, files[1], duplicates.items[1].files[0])
	assertEqual(t, files[3], duplicates.items[1].files[1])
}

func TestNewDuplicatesStats(t *testing.T) {
	files := []*FileInfo{
		{path: "root/subdir1/a.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir1/b.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}, size: 2},
		{path: "root/subdir2/c.txt", hash: []byte{0x3, 0x4, 0x5, 0x6}, size: 3},
		{path: "root/subdir2/d.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}, size: 2},
		{path: "root/subdir3/e.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir3/f.txt", hash: []byte{0x4, 0x5, 0x6, 0x7}, size: 5},
		{path: "root/subdir4/g.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir4/h.txt", hash: []byte{0x5, 0x6, 0x7, 0x8}, size: 1},
	}

	dirTree := DirTree{rootPath: "root", files: files}
	duplicates := NewDuplicates(&dirTree)

	stats := duplicates.Stats()

	assertEqual(t, 8, stats.num_files_total);
	assertEqual(t, 5, stats.num_duplicate_files);
	assertEqual(t, 16, stats.size_files_total);
	assertEqual(t,  7, stats.size_duplicate_files);
}

func TestNewDuplicatesWithSameHashButDifferenzFileSize(t *testing.T) {
	files := []*FileInfo{
		{path: "root/subdir1/a.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 1},
		{path: "root/subdir1/b.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 2},
		{path: "root/subdir2/c.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 3},
		{path: "root/subdir2/d.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 4},
		{path: "root/subdir3/e.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 5},
		{path: "root/subdir3/f.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 6},
		{path: "root/subdir4/g.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 7},
		{path: "root/subdir4/h.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}, size: 8},
	}

	dirTree := DirTree{rootPath: "root", files: files}

	duplicates := NewDuplicates(&dirTree)

	assertEqual(t, 0, duplicates.Count())
}
