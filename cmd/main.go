package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/somaz94/go-docs-generator-action/pkg/template"
)

type Config struct {
	SourcePath   string
	OutputPath   string
	TemplateType string
	TemplateName string
	Format       string
}

func main() {
	config := Config{
		SourcePath:   getEnvWithDefault("INPUT_SOURCE_PATH", "."),
		OutputPath:   getEnvWithDefault("INPUT_OUTPUT_PATH", "docs"),
		TemplateType: getEnvWithDefault("INPUT_TEMPLATE_TYPE", "readme"),
		TemplateName: getEnvWithDefault("INPUT_TEMPLATE_NAME", "default"),
		Format:       getEnvWithDefault("INPUT_FORMAT", "markdown"),
	}

	if err := run(config); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func run(config Config) error {
	// Initialize template manager
	tm := template.NewTemplateManager("templates")

	// Load template
	if err := tm.LoadTemplate(config.TemplateType, config.TemplateName); err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// Prepare example data (실제로는 프로젝트에서 데이터를 수집해야 함)
	data := &template.ReadmeData{
		ProjectName: "Go Docs Generator Action",
		Description: "Automatically generate documentation for your project",
		Features: []string{
			"README generation",
			"API documentation",
			"Example code documentation",
		},
		Requirements:      "Go 1.21 or higher",
		InstallationSteps: "go get github.com/yourusername/go-docs-generator-action",
		Language:          "go",
		UsageExample: `func main() {
    fmt.Println("Hello, Documentation!")
}`,
		ConfigOptions: []template.ConfigOption{
			{
				Name:        "source_path",
				Type:        "string",
				Default:     ".",
				Description: "Source code path",
			},
		},
		License: "MIT",
		Contributors: []template.Contributor{
			{
				Name: "somaz94",
				Role: "Maintainer",
			},
		},
	}

	// Execute template
	result, err := tm.Execute(config.TemplateType, config.TemplateName, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Write output
	outputPath := filepath.Join(config.OutputPath, "README.md")
	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Printf("Documentation generated successfully: %s\n", outputPath)
	return nil
}
