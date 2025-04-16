// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/washiil/euler-go/problems/p001"
	_ "github.com/washiil/euler-go/problems/p002"
	_ "github.com/washiil/euler-go/problems/p003"
	_ "github.com/washiil/euler-go/problems/p004"
	_ "github.com/washiil/euler-go/problems/p005"
	_ "github.com/washiil/euler-go/problems/p006"
	_ "github.com/washiil/euler-go/problems/p007"
	_ "github.com/washiil/euler-go/problems/p008"
	_ "github.com/washiil/euler-go/problems/p009"
	_ "github.com/washiil/euler-go/problems/p010"
	_ "github.com/washiil/euler-go/problems/p011"
	_ "github.com/washiil/euler-go/problems/p012"
	_ "github.com/washiil/euler-go/problems/p013"

	"github.com/washiil/euler-go/problems/registry"
)

func main() {
	register := registry.Problems

	if len(os.Args) < 2 {
		fmt.Printf("Please input a valid problem number\n$ go run main <number>\n")
		return
	} else if len(os.Args) == 3 && os.Args[1] == "create" {
		// Creating file
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid problem number please try again.")
			return
		}
		createFromTemplate(num)
		return
	}

	arg := os.Args[1]
	if fn, ok := register[arg]; ok {
		result := fn()
		fmt.Printf("Problem [%s] : %d", arg, result)
	} else if arg == "all" {
		for key, fn := range register {
			result := fn()
			fmt.Printf("Problem [%s] : %d\n", key, result)
		}
	} else if arg == "benchmark" {
		if len(os.Args) < 3 {
			fmt.Printf("Please input a valid problem number\n$ go run main benchmark <number>\n")
			return
		}
		if fn, ok := register[os.Args[2]]; ok {
			const samples = 10
			fmt.Printf("Sampling function %d times\n", samples)
			avg_runtime := benchmark(fn, samples)
			fmt.Printf("Problem [%s] : %s\n", os.Args[2], avg_runtime)
		}
	} else {
		fmt.Printf("Problem %s not implemented.\n", arg)
	}
}

func benchmark(function func() int, samples int) time.Duration {
	var total_duration time.Duration

	for i := 0; i < samples; i++ {
		start := time.Now()
		function()
		elapsed := time.Since(start)
		total_duration += elapsed
	}

	average := total_duration / time.Duration(samples)
	return average
}

func createFromTemplate(number int) error {
	replacement := fmt.Sprintf("%03d", number)
	templateFile := "template.txt"
	outputFile := fmt.Sprintf("problems/p%s/%s.go", replacement, replacement)
	placeholder := "__PROBLEM_NUMBER__"

	// 1. Read the entire content of the template file
	templateBytes, err := os.ReadFile(templateFile)
	if err != nil {
		// Wrap the error with more context
		return fmt.Errorf("failed to read template file '%s': %w", templateFile, err)
	}
	templateContent := string(templateBytes)

	// 2. Perform the replacement
	// strings.ReplaceAll finds *all* non-overlapping instances of 'placeholder'
	// and replaces them with 'replacementValue'.
	outputContent := strings.ReplaceAll(templateContent, placeholder, replacement)

	// 3. Ensure the output directory exists (optional but good practice)
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil { // 0755 permissions: rwxr-xr-x
		return fmt.Errorf("failed to create output directory '%s': %w", outputDir, err)
	}

	// 4. Write the modified content to the new output file
	// os.WriteFile creates the file if it doesn't exist, or truncates it if it does.
	// 0644 permissions are common for files (rw-r--r--).
	err = os.WriteFile(outputFile, []byte(outputContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file '%s': %w", outputFile, err)
	}

	// If we reached here, everything was successful
	fmt.Printf("Successfully created '%s' from template '%s'\n", outputFile, templateFile)
	return nil
}
