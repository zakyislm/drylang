package cli

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/*
var templateFiles embed.FS

func handleInit(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: dry init <folder> [template]")
		os.Exit(1)
	}

	targetDir := args[0]
	if targetDir != "." {
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			die(fmt.Errorf("failed to create directory: %v", err))
		}
	}

	targetFile := filepath.Join(targetDir, "idx.dry")
	if _, err := os.Stat(targetFile); err == nil {
		fmt.Printf("Error: %s already exists\n", targetFile)
		os.Exit(1)
	}

	var templateName string
	if len(args) >= 2 {
		templateName = args[1]
	}

	templateMap := map[string]string{
		"api":        "rest-api.dry",
		"web":        "file-server.dry",
		"crud":       "crud.dry",
		"fetch":      "fetch-json.dry",
		"cli":        "cli-tool.dry",
		"automation": "automation.dry",
		"scraper":    "scraper.dry",
		"hello":      "hello.dry",
	}

	var content []byte
	if templateName == "" {
		content = []byte("// idx.dry - Entry point\npt \"Hello, dryLang!\"\n")
	} else {
		fileName, ok := templateMap[templateName]
		if !ok {
			fmt.Printf("Unknown template: %s\n", templateName)
			fmt.Println("Available templates: api, web, crud, fetch, cli, automation, scraper, hello")
			os.Exit(1)
		}

		var err error
		content, err = templateFiles.ReadFile("templates/" + fileName)
		if err != nil {
			die(fmt.Errorf("failed to load template: %v", err))
		}
	}

	if err := os.WriteFile(targetFile, content, 0644); err != nil {
		die(fmt.Errorf("failed to write file: %v", err))
	}

	fmt.Printf("Successfully created %s", targetFile)
	if templateName != "" {
		fmt.Printf(" using template '%s'", templateName)
	}
	fmt.Println("\nRun it with: dry " + targetDir)
}
