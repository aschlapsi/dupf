package dupfinder

type Duplicates struct {
	Items []*DuplicateItems
}

type DuplicateItems struct {
	Key string
	Files []*FileInfo
}

func NewDuplicates(dirTree *DirTree) *Duplicates {
	duplicates := &Duplicates{Items: make([]*DuplicateItems, 0)}
	fileMap := make(map[string][]*FileInfo)

	for _, fi := range dirTree.Files {
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
			duplicates.Items = append(duplicates.Items, &DuplicateItems{Key: key, Files: fileList})
		}
	}

	return duplicates
}

func (dup *Duplicates) Count() int {
	return len(dup.Items)
}