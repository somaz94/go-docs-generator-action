package template

import (
	"testing"
)

func TestTemplateManager(t *testing.T) {
	tm := NewTemplateManager("../../templates")

	// Test template loading
	err := tm.LoadTemplate("readme", "default")
	if err != nil {
		t.Fatalf("Failed to load template: %v", err)
	}

	// Test template execution
	data := &ReadmeData{
		ProjectName: "Test Project",
		Description: "Test Description",
	}

	result, err := tm.Execute("readme", "default", data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	if result == "" {
		t.Error("Template execution resulted in empty string")
	}
}
