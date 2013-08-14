package dupfinder

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirTree struct {
	rootPath string
	files []*FileInfo
}

func NewDirTree(rootDir string) *DirTree {
	return &DirTree{rootPath: rootDir, files: walkDirectory(rootDir)}
}

func (dt *DirTree) FileCount() int {
	return len((*dt).files)
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
	path string
	hash []byte
}

func NewFileInfo(path string) *FileInfo {
	return &FileInfo{path: path, hash: getHash(path)}
}

func (fi *FileInfo) GetHashstring() string {
	return fmt.Sprintf("%x", fi.hash)
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
