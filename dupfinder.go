package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type SearchProgress struct {
	progress chan string
	files []*FileInfo
	totalFileSize int64
	duplicateMap map[string][]FileInfo
	duplicates *Duplicates
}

func (sp *SearchProgress) Progress() <-chan string {
	return sp.progress
}

func (sp *SearchProgress) TotalFileSize() int64 {
	return sp.totalFileSize
}

func (sp *SearchProgress) FileCount() int {
	return len(sp.files)
}

func (sp *SearchProgress) Duplicates() *Duplicates {
	if sp.duplicates == nil {
		sp.duplicates = &Duplicates{duplicateGroups: make([][]FileInfo, 0)}

		for _, v := range sp.duplicateMap {
			if len(v) > 1 {
				for _, fi := range v {
					sp.duplicates.totalDuplicateFileSize += fi.size
				}
				sp.duplicates.duplicateGroups = append(sp.duplicates.duplicateGroups, v)
			}
		}
	}

	return sp.duplicates
}

func (sp *SearchProgress) walkFn(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		absPath, err := filepath.Abs(path)
		if err != nil {
			absPath = path
		}
		sp.progress <- absPath
		sp.addFileInfo(absPath)
	}

	return nil
}

func (sp *SearchProgress) addFileInfo(path string) {
	hash := getHash(path)
	size := getSize(path)
	sp.totalFileSize += size
	fileInfo := FileInfo{path: path, hash: hash, size: size}

	if sp.duplicateMap == nil {
		sp.duplicateMap = make(map[string][]FileInfo)
	}

	key := fileInfo.Hashstring()
	sp.duplicateMap[key] = append(sp.duplicateMap[key], fileInfo)

	sp.files = append(sp.files, &fileInfo)
}

func (sp *SearchProgress) close() {
	close(sp.progress)
}

type Duplicates struct {
	totalDuplicateFileSize int64
	duplicateGroups [][]FileInfo
}

func (d Duplicates) Count() int {
	return len(d.duplicateGroups)
}

func (d Duplicates) TotalFileSize() int64 {
	return d.totalDuplicateFileSize
}

func (d Duplicates) Groups() [][]FileInfo {
	return d.duplicateGroups
}

type FileInfo struct {
	path string
	hash []byte
	size int64
}

func (fi *FileInfo) Hashstring() string {
	return fmt.Sprintf("%x", fi.hash)
}

func (fi *FileInfo) Path() string {
	return fi.path
}

func FindDuplicates(dirs ...string) *SearchProgress {
	progress := make(chan string)
	search := &SearchProgress{progress: progress}

	go func() {
		for _, dir := range dirs {
			filepath.Walk(dir, search.walkFn)
		}

		search.close()
	}()

	return search
}

func getSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}

	return fileInfo.Size()
}

var hash = md5.New()

func getHash(path string) []byte {
	hash.Reset()

	// TODO: generate hash using smaller byte blocks

	content, err := ioutil.ReadFile(path)
	if err != nil {
		// TODO: better error handling
		fmt.Printf("Error while reading file %s: %s\n", path, err)
		return nil
	}

	_, err = hash.Write(content)
	if err != nil {
		// TODO: better error handling
		fmt.Printf("Error while generating hash for file %s: %s\n", path, err)
		return nil
	}

	return hash.Sum(nil)
}