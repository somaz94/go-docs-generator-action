package template

// ReadmeData represents the data structure for README template
type ReadmeData struct {
	ProjectName       string
	Description       string
	Features          []string
	Requirements      string
	InstallationSteps string
	Language          string
	UsageExample      string
	ConfigOptions     []ConfigOption
	License           string
	Contributors      []Contributor
}

// ConfigOption represents a configuration option
type ConfigOption struct {
	Name        string
	Type        string
	Default     string
	Description string
}

// Contributor represents a project contributor
type Contributor struct {
	Name string
	Role string
}
