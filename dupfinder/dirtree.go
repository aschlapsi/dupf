package dupfinder

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirTree struct {
	RootPath string
	Files []*FileInfo
}

func NewDirTree(rootDir string) *DirTree {
	return &DirTree{RootPath: rootDir, Files: walkDirectory(rootDir)}
}

func (dt *DirTree) FileCount() int {
	return len((*dt).Files)
}

func walkDirectory(rootDir string) (result []*FileInfo) {
	walker := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			result = append(result, NewFileInfo(path))
		}

		return nil
	}

	filepath.Walk(rootDir, walker)
	return
}

type FileInfo struct {
	Path string
	Hash []byte
}

func NewFileInfo(path string) *FileInfo {
	return &FileInfo{Path: path, Hash: getHash(path)}
}

func (fi *FileInfo) GetHashstring() string {
	return fmt.Sprintf("%x", fi.Hash)
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
