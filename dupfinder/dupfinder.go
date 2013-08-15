package dupfinder

import "strconv"

type Duplicates struct {
	items []*DuplicateItems
	stats *Stats
}

type DuplicateItems struct {
	key string
	files []*FileInfo
}

type Stats struct {
	num_files_total int
	num_duplicate_files int
	size_files_total int64
	size_duplicate_files int64
}

func NewDuplicates(dirTree *DirTree) *Duplicates {
	stats := &Stats{}
	duplicates := &Duplicates{items: make([]*DuplicateItems, 0), stats: stats}
	fileMap := make(map[string][]*FileInfo)

	for _, fi := range dirTree.files {
		stats.num_files_total++
		stats.size_files_total += fi.size
		hashstr := fi.GetHashstring()
		fileKey := hashstr + strconv.FormatInt(fi.size, 10)

		files, ok := fileMap[fileKey]
		if !ok {
			files = make([]*FileInfo, 0)
			fileMap[fileKey] = files
		}
		fileMap[fileKey] = append(files, fi)
	}

	for key, fileList := range fileMap {
		if len(fileList) > 1 {
			stats.num_duplicate_files += len(fileList)
			stats.size_duplicate_files += int64(len(fileList)) * fileList[0].size
			duplicates.items = append(duplicates.items, &DuplicateItems{key: key, files: fileList})
		}
	}

	return duplicates
}

func (dup *Duplicates) Count() int {
	return len(dup.items)
}

func (dup *Duplicates) Stats() *Stats {
	return dup.stats
}