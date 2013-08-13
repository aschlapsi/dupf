package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Hash []byte
}

func (fi *FileInfo) GetHashstring() string {
	return fmt.Sprintf("%x", fi.Hash)
}

func WalkDirectory(rootDir string) (result []FileInfo) {
	walker := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			result = append(result, newFileInfo(path))
		}

		return nil
	}

	filepath.Walk(rootDir, walker)
	return
}

func newFileInfo(path string) FileInfo {
	return FileInfo{Path: path, Hash: getHash(path)}
}

var hash = md5.New()

func getHash(path string) []byte {
	hash.Reset()

	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error while reading file %s: %s\n", path, err)
		return nil
	}

	_, err = hash.Write(content)
	if err != nil {
		fmt.Printf("Error while generating hash for file %s: %s\n", path, err)
		return nil
	}

	return hash.Sum(nil)
}
