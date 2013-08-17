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
	duplicateMap map[string]*FileInfo
}

func (sp *SearchProgress) Progress() <-chan string {
	return sp.progress
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
	sp.files = append(sp.files, &FileInfo{path: path, hash: hash})
}

func (sp *SearchProgress) close() {
	close(sp.progress)
}

type FileInfo struct {
	path string
	hash []byte
	size int64
}

func (fi *FileInfo) Hashstring() string {
	return fmt.Sprintf("%x", fi.hash)
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