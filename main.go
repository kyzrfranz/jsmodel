package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kyzrfranz/jsmodel/pkg/generator"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	urlPtr := flag.String("url", "", "URL of the JSON Schema")
	filePtr := flag.String("file", "", "Path to the local JSON Schema file")
	outputPtr := flag.String("o", "", "Path to the output folder for the generated file(s) (optional)")
	flag.Parse()

	var fileData []byte
	var err error

	switch {
	case *urlPtr != "":
		// Fetch JSON schema from URL
		resp, err := http.Get(*urlPtr)
		if err != nil {
			fmt.Println("Error fetching JSON schema:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fileData, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}
	case *filePtr != "":
		// Read JSON schema from file
		fileData, err = os.ReadFile(*filePtr)
		if err != nil {
			fmt.Println("Error reading JSON schema file:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Please provide a URL or a file path for the JSON Schema")
		os.Exit(1)
	}

	// Parse JSON schema
	var schema generator.JSONSchema
	err = json.Unmarshal(fileData, &schema)
	if err != nil {
		fmt.Println("Error parsing JSON schema:", err)
		os.Exit(1)
	}

	// Generate JavaScript class from schema
	for className, def := range schema.Definitions {
		jsClass := generator.GenerateJSClass(className, def)

		// Output to file or stdout
		if *outputPtr != "" {
			outputPath := filepath.Join(*outputPtr, fmt.Sprintf("%s.js", className))
			err = os.WriteFile(outputPath, []byte(jsClass), 0644)
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				os.Exit(1)
			}
			fmt.Printf("Generated JavaScript class written to: %s\n", outputPath)
		} else {
			fmt.Println(jsClass)
		}
	}
}
