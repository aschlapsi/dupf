package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: dupf <directory>...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	search := FindDuplicates(args...)

	for {
		filePath := <-search.Progress()
		if filePath == "" {
			break;
		}

		fmt.Printf("\033[1K\rProcessing file %s.", filePath)
	}

	fmt.Println()
}


/*
	var dirTree *dupfinder.DirTree

	for _, dir := range args {
		if dirTree == nil {
			dirTree = dupfinder.NewDirTree(dir)
		} else {
			dirTree.AppendDir(dir)
		}
	}

	duplicates := dupfinder.NewDuplicates(dirTree)
	printDuplicateInfo(duplicates)
}

func printDuplicateInfo(duplicates *dupfinder.Duplicates) {
	printStats(duplicates.Stats())
}

func printStats(stats *dupfinder.Stats) {
	fmt.Printf("Processed %d files and found %d duplicates.\n", stats.TotalFiles(), stats.DuplicateFiles())
	fmt.Printf("The total file size was %d bytes and the duplicate files had a size of %d bytes.\n", stats.TotalFileSize(), stats.DuplicateFilesSize())
	fmt.Printf("Without duplicates the total file size would be %d bytes.\n", stats.TotalFileSizeWithoutDuplicates())
}
*/