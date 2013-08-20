package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: dupf <directory>...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	start := time.Now()
	search := FindDuplicates(args...)

	for {
		filePath := <-search.Progress()
		if filePath == "" {
			break;
		}

		fmt.Printf("\033[1K\rProcessing file %s", filePath)
	}

	duration := time.Since(start)
	showResult(search, duration)
}

func showResult(search *SearchProgress, duration time.Duration) {
	duplicates := search.Duplicates()

	fmt.Println()
	fmt.Println()
	fmt.Printf("Found %d duplicate groups:\n", duplicates.Count())
	for i, group := range duplicates.Groups() {
		fmt.Printf("    - Group #%d\n", i + 1)
		for _, file := range group {
			fmt.Printf("          - %s\n", file.Path())
		}
	}
	fmt.Println()
	fmt.Printf("Processed %d files with a total file size of %d bytes in %s.\n", search.FileCount(), search.TotalFileSize(), duration.String())
	fmt.Printf("Total size of duplicated files: %d bytes\n", duplicates.TotalFileSize())
	fmt.Println()
}
