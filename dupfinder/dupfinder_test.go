package dupfinder

import (
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	files := []*FileInfo{
		{path: "root/subdir1/a.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{path: "root/subdir1/b.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}},
		{path: "root/subdir2/c.txt", hash: []byte{0x3, 0x4, 0x5, 0x6}},
		{path: "root/subdir2/d.txt", hash: []byte{0x2, 0x3, 0x4, 0x5}},
		{path: "root/subdir3/e.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{path: "root/subdir3/f.txt", hash: []byte{0x4, 0x5, 0x6, 0x7}},
		{path: "root/subdir4/g.txt", hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{path: "root/subdir4/h.txt", hash: []byte{0x5, 0x6, 0x7, 0x8}},
	}

	dirTree := DirTree{rootPath: "root", files: files}

	duplicates := NewDuplicates(&dirTree)

	assertEqual(t, 2, duplicates.Count())
	assertEqual(t, "01020304", duplicates.items[0].key)
	assertEqual(t, 3, len(duplicates.items[0].files))
	assertEqual(t, files[0], duplicates.items[0].files[0])
	assertEqual(t, files[4], duplicates.items[0].files[1])
	assertEqual(t, files[6], duplicates.items[0].files[2])
	assertEqual(t, "02030405", duplicates.items[1].key)
	assertEqual(t, 2, len(duplicates.items[1].files))
	assertEqual(t, files[1], duplicates.items[1].files[0])
	assertEqual(t, files[3], duplicates.items[1].files[1])
}