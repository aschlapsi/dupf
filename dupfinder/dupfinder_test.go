package dupfinder

import (
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	files := []*FileInfo{
		{Path: "root/subdir1/a.txt", Hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{Path: "root/subdir1/b.txt", Hash: []byte{0x2, 0x3, 0x4, 0x5}},
		{Path: "root/subdir2/c.txt", Hash: []byte{0x3, 0x4, 0x5, 0x6}},
		{Path: "root/subdir2/d.txt", Hash: []byte{0x2, 0x3, 0x4, 0x5}},
		{Path: "root/subdir3/e.txt", Hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{Path: "root/subdir3/f.txt", Hash: []byte{0x4, 0x5, 0x6, 0x7}},
		{Path: "root/subdir4/g.txt", Hash: []byte{0x1, 0x2, 0x3, 0x4}},
		{Path: "root/subdir4/h.txt", Hash: []byte{0x5, 0x6, 0x7, 0x8}},
	}

	dirTree := DirTree{RootPath: "root", Files: files}

	duplicates := NewDuplicates(&dirTree)

	assertEqual(t, 2, duplicates.Count())
	assertEqual(t, "01020304", duplicates.Items[0].Key)
	assertEqual(t, 3, len(duplicates.Items[0].Files))
	assertEqual(t, files[0], duplicates.Items[0].Files[0])
	assertEqual(t, files[4], duplicates.Items[0].Files[1])
	assertEqual(t, files[6], duplicates.Items[0].Files[2])
	assertEqual(t, "02030405", duplicates.Items[1].Key)
	assertEqual(t, 2, len(duplicates.Items[1].Files))
	assertEqual(t, files[1], duplicates.Items[1].Files[0])
	assertEqual(t, files[3], duplicates.Items[1].Files[1])
}