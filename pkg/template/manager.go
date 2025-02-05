package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// TemplateManager handles loading and executing templates
type TemplateManager struct {
	templates map[string]*template.Template
	basePath  string
}

// NewTemplateManager creates a new template manager
func NewTemplateManager(basePath string) *TemplateManager {
	return &TemplateManager{
		templates: make(map[string]*template.Template),
		basePath:  basePath,
	}
}

// LoadTemplate loads a template from the filesystem
func (tm *TemplateManager) LoadTemplate(category, name string) error {
	path := filepath.Join(tm.basePath, category, name+".md")
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read template %s/%s: %w", category, name, err)
	}

	// Create template with custom functions
	tmpl, err := template.New(name).Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
		"join":    strings.Join,
	}).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s/%s: %w", category, name, err)
	}

	tm.templates[category+"/"+name] = tmpl
	return nil
}

// Execute applies the template with given data
func (tm *TemplateManager) Execute(category, name string, data interface{}) (string, error) {
	tmplKey := category + "/" + name
	tmpl, ok := tm.templates[tmplKey]
	if !ok {
		return "", fmt.Errorf("template not found: %s", tmplKey)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", tmplKey, err)
	}

	return buf.String(), nil
}

// ListTemplates returns all available templates
func (tm *TemplateManager) ListTemplates() []string {
	var templates []string
	for key := range tm.templates {
		templates = append(templates, key)
	}
	return templates
}

// ValidateTemplate checks if a template exists
func (tm *TemplateManager) ValidateTemplate(category, name string) bool {
	_, ok := tm.templates[category+"/"+name]
	return ok
}
