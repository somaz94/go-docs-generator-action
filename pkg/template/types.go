package template

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// // ReadmeData represents the data structure for simple README template
// type ReadmeData struct {
// 	ProjectName     string
// 	GitHubRepo      string
// 	MarketplaceName string
// 	MarketplaceSlug string
// 	Description     string
// 	Features        []string
// 	License         string
// 	Contributing    string
// }

// DetailedReadmeData represents the data structure for detailed README template
type DetailedReadmeData struct {
	ProjectName       string
	GitHubRepo        string
	MarketplaceName   string
	MarketplaceSlug   string
	Description       string
	Features          []string
	Requirements      string
	InstallationSteps string
	Language          string
	UsageExample      string
	ConfigOptions     []ConfigOption
	License           string
	Contributors      []Contributor
	Contributing      string
	ActionConfig      bool
}

// ConfigOption represents a configuration option
type ConfigOption struct {
	Name        string
	Required    bool
	Description string
	Default     string
}

// Contributor represents a project contributor
type Contributor struct {
	Name string
	Role string
}

// GetActionInputs reads action.yml and returns config options
func GetActionInputs() ([]ConfigOption, error) {
	data, err := os.ReadFile("action.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to read action.yml: %w", err)
	}

	var action struct {
		Inputs map[string]struct {
			Description string `yaml:"description"`
			Required    bool   `yaml:"required"`
			Default     string `yaml:"default"`
		} `yaml:"inputs"`
	}

	if err := yaml.Unmarshal(data, &action); err != nil {
		return nil, fmt.Errorf("failed to parse action.yml: %w", err)
	}

	var options []ConfigOption
	for name, input := range action.Inputs {
		options = append(options, ConfigOption{
			Name:        name,
			Required:    input.Required,
			Description: input.Description,
			Default:     input.Default,
		})
	}

	return options, nil
}
