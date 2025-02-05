package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTemplateManager(t *testing.T) {
	// Setup test template directory and cleanup
	tmpDir := setupTestTemplates(t)
	defer os.RemoveAll(tmpDir)

	tm := NewTemplateManager(tmpDir)

	t.Run("Template Loading", func(t *testing.T) {
		err := tm.LoadTemplate("readme", "default")
		if err != nil {
			t.Fatalf("Failed to load template: %v", err)
		}

		// Test invalid template
		err = tm.LoadTemplate("invalid", "invalid")
		if err == nil {
			t.Error("Expected error for invalid template, got nil")
		}
	})

	t.Run("Template Execution", func(t *testing.T) {
		data := &ReadmeData{
			ProjectName:     "Test Project",
			GitHubRepo:      "test/project",
			MarketplaceName: "Test Action",
			MarketplaceSlug: "test-action",
			License:         "MIT",
			Contributing:    "Test contributing message",
		}

		result, err := tm.Execute("readme", "default", data)
		if err != nil {
			t.Fatalf("Failed to execute template: %v", err)
		}

		// Check if result contains expected content
		expectedContents := []string{
			"# Test Project",
			"[![License]",
			"test/project",
			"MIT",
			"Test contributing message",
		}

		for _, expected := range expectedContents {
			if !strings.Contains(result, expected) {
				t.Errorf("Expected result to contain %q", expected)
			}
		}
	})

	t.Run("Template Validation", func(t *testing.T) {
		if !tm.ValidateTemplate("readme", "default") {
			t.Error("Expected template to be valid")
		}

		if tm.ValidateTemplate("invalid", "invalid") {
			t.Error("Expected template to be invalid")
		}
	})

	t.Run("Template Listing", func(t *testing.T) {
		templates := tm.ListTemplates()
		if len(templates) != 1 {
			t.Errorf("Expected 1 template, got %d", len(templates))
		}
		if templates[0] != "readme/default" {
			t.Errorf("Expected template name 'readme/default', got %q", templates[0])
		}
	})
}

// setupTestTemplates creates a temporary directory with test templates
func setupTestTemplates(t *testing.T) string {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "template-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	readmeDir := filepath.Join(tmpDir, "readme")
	if err := os.MkdirAll(readmeDir, 0755); err != nil {
		t.Fatalf("Failed to create readme dir: %v", err)
	}

	// Create test template
	template := `# {{ .ProjectName }}

[![License](https://img.shields.io/github/license/{{ .GitHubRepo }})](https://github.com/{{ .GitHubRepo }})
![Latest Tag](https://img.shields.io/github/v/tag/{{ .GitHubRepo }})
![Top Language](https://img.shields.io/github/languages/top/{{ .GitHubRepo }}?color=green&logo=go&logoColor=b)
[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-{{ .MarketplaceName }}-blue?logo=github)](https://github.com/marketplace/actions/{{ .MarketplaceSlug }})

## License

This project is licensed under the {{ .License }} License - see the [LICENSE](LICENSE) file for details.

## Contributing

{{ .Contributing }}`

	if err := os.WriteFile(filepath.Join(readmeDir, "default.md"), []byte(template), 0644); err != nil {
		t.Fatalf("Failed to write template file: %v", err)
	}

	return tmpDir
}
