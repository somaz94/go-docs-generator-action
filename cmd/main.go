package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/somaz94/go-docs-generator-action/pkg/template"
)

type Config struct {
	OutputPath          string
	TemplateType        string
	TemplateNames       []string
	Format              string
	ProjectName         string
	GitHubRepo          string
	MarketplaceName     string
	MarketplaceSlug     string
	Description         string
	Features            []string
	Requirements        string
	InstallationSteps   string
	UsageExample        string
	Language            string
	License             string
	ContributingMessage string
	ConfigOptions       []template.ConfigOption
	Contributors        []template.Contributor
	ActionConfig        bool
}

func main() {
	templateNames := strings.Split(getEnvWithDefault("INPUT_TEMPLATE_NAMES", "default"), ",")
	features := strings.Split(getEnvWithDefault("INPUT_FEATURES", ""), ",")
	if features[0] == "" {
		features = nil
	}

	// ConfigOptions 파싱 (예: "name:type:default:description,name2:type2:default2:description2")
	var configOptions []template.ConfigOption
	configOptionsStr := getEnvWithDefault("INPUT_CONFIG_OPTIONS", "")
	if configOptionsStr != "" {
		for _, opt := range strings.Split(configOptionsStr, ",") {
			parts := strings.Split(opt, ":")
			if len(parts) == 4 {
				configOptions = append(configOptions, template.ConfigOption{
					Name:        parts[0],
					Type:        parts[1],
					Default:     parts[2],
					Description: parts[3],
				})
			}
		}
	}

	// Contributors 파싱 (예: "name:role,name2:role2")
	var contributors []template.Contributor
	contributorsStr := getEnvWithDefault("INPUT_CONTRIBUTORS", "")
	if contributorsStr != "" {
		for _, cont := range strings.Split(contributorsStr, ",") {
			parts := strings.Split(cont, ":")
			if len(parts) == 2 {
				contributors = append(contributors, template.Contributor{
					Name: parts[0],
					Role: parts[1],
				})
			}
		}
	}

	config := Config{
		OutputPath:          getEnvWithDefault("INPUT_OUTPUT_PATH", "docs"),
		TemplateType:        getEnvWithDefault("INPUT_TEMPLATE_TYPE", "readme"),
		TemplateNames:       templateNames,
		Format:              getEnvWithDefault("INPUT_FORMAT", "markdown"),
		ProjectName:         os.Getenv("INPUT_PROJECT_NAME"),
		GitHubRepo:          os.Getenv("INPUT_GITHUB_REPO"),
		MarketplaceName:     os.Getenv("INPUT_MARKETPLACE_NAME"),
		MarketplaceSlug:     os.Getenv("INPUT_MARKETPLACE_SLUG"),
		Description:         getEnvWithDefault("INPUT_DESCRIPTION", ""),
		Features:            features,
		Requirements:        getEnvWithDefault("INPUT_REQUIREMENTS", ""),
		InstallationSteps:   getEnvWithDefault("INPUT_INSTALLATION_STEPS", ""),
		UsageExample:        getEnvWithDefault("INPUT_USAGE_EXAMPLE", ""),
		Language:            getEnvWithDefault("INPUT_LANGUAGE", "go"),
		License:             getEnvWithDefault("INPUT_LICENSE", "MIT"),
		ContributingMessage: getEnvWithDefault("INPUT_CONTRIBUTING_MESSAGE", "Contributions are welcome! Please feel free to submit a Pull Request"),
		ConfigOptions:       configOptions,
		Contributors:        contributors,
		ActionConfig:        getEnvWithDefault("INPUT_ACTION_CONFIG", "false") == "true",
	}

	if config.ProjectName == "" || config.GitHubRepo == "" || config.MarketplaceName == "" || config.MarketplaceSlug == "" {
		log.Fatal("Required inputs are missing")
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
	tm := template.NewTemplateManager("templates")

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(config.OutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if config.ActionConfig {
		actionInputs, err := template.GetActionInputs()
		if err != nil {
			log.Printf("Warning: failed to get action inputs: %v", err)
		} else {
			config.ConfigOptions = append(config.ConfigOptions, actionInputs...)
		}
	}

	// Process each template
	for _, templateName := range config.TemplateNames {
		// Load template
		if err := tm.LoadTemplate(config.TemplateType, templateName); err != nil {
			return fmt.Errorf("failed to load template %s: %w", templateName, err)
		}

		// Prepare data based on template type
		var data interface{}
		switch templateName {
		case "default", "detailed":
			data = &template.DetailedReadmeData{
				ProjectName:       config.ProjectName,
				Description:       config.Description,
				Features:          config.Features,
				Requirements:      config.Requirements,
				InstallationSteps: config.InstallationSteps,
				Language:          config.Language,
				UsageExample:      config.UsageExample,
				ConfigOptions:     config.ConfigOptions,
				License:           config.License,
				Contributors:      config.Contributors,
			}
		}

		// Execute template
		result, err := tm.Execute(config.TemplateType, templateName, data)
		if err != nil {
			return fmt.Errorf("failed to execute template %s: %w", templateName, err)
		}

		// Write output
		outputFile := "README.md"
		if templateName != "default" {
			outputFile = fmt.Sprintf("README-%s.md", templateName)
		}
		outputPath := filepath.Join(config.OutputPath, outputFile)
		if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
			return fmt.Errorf("failed to write output for %s: %w", templateName, err)
		}

		fmt.Printf("Documentation generated successfully: %s\n", outputPath)
	}

	return nil
}
