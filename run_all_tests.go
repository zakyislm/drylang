package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func main() {
	files, err := filepath.Glob("test/level*_test*.y")
	if err != nil {
		fmt.Println("Error globbing:", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("No test files found.")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	errCh := make(chan string, len(files))

	// Limit concurrency
	sem := make(chan struct{}, 8)

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			cmd := exec.Command("go", "run", ".", "run", f)
			output, err := cmd.CombinedOutput()
			if err != nil {
				errCh <- fmt.Sprintf("FAIL: %s\nOutput:\n%s\nError: %v\n", f, string(output), err)
			}
		}(file)
	}

	wg.Wait()
	close(errCh)

	failed := false
	for errStr := range errCh {
		fmt.Println(errStr)
		failed = true
	}

	if !failed {
		fmt.Printf("ALL %d TESTS PASSED!\n", len(files))
	} else {
		fmt.Println("SOME TESTS FAILED.")
		os.Exit(1)
	}
}
