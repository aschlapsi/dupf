package main

import (
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
}

func walkDirectory(rootDir string) (result []FileInfo) {
	walker := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			result = append(result, FileInfo{Path: path})
		}

		return nil
	}

	filepath.Walk(rootDir, walker)
	return
}