package dupfinder

type Duplicates struct {
	items []*DuplicateItems
}

type DuplicateItems struct {
	key string
	files []*FileInfo
}

func NewDuplicates(dirTree *DirTree) *Duplicates {
	duplicates := &Duplicates{items: make([]*DuplicateItems, 0)}
	fileMap := make(map[string][]*FileInfo)

	for _, fi := range dirTree.files {
		hashstr := fi.GetHashstring()

		files, ok := fileMap[hashstr]
		if !ok {
			files = make([]*FileInfo, 0)
			fileMap[hashstr] = files
		}
		fileMap[hashstr] = append(files, fi)
	}

	for key, fileList := range fileMap {
		if len(fileList) > 1 {
			duplicates.items = append(duplicates.items, &DuplicateItems{key: key, files: fileList})
		}
	}

	return duplicates
}

func (dup *Duplicates) Count() int {
	return len(dup.items)
}