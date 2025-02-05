package template

// ReadmeData represents the data structure for simple README template
type ReadmeData struct {
	ProjectName     string
	GitHubRepo      string
	MarketplaceName string
	MarketplaceSlug string
	License         string
	Contributing    string
}

// DetailedReadmeData represents the data structure for detailed README template
type DetailedReadmeData struct {
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
