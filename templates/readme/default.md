# {{ .ProjectName }}

# Overview

{{ .Description }}

## Inputs

| Input | Required | Description | Default |
|-------|----------|-------------|---------|
{{ range .ConfigOptions }}
| `{{ .Name }}` | {{ .Required }} | {{ .Description }} | {{ .Default }} |
{{ end }}

## 🚀 Features

{{ range .Features }}
- {{ . }}
{{ end }}

## 📋 Requirements

{{ .Requirements }}

## 🔧 Installation

```bash
{{ .InstallationSteps }}
```

## 💡 Usage

```{{ .Language }}
{{ .UsageExample }}
```

## 📝 License

{{ .License }}

## 👥 Contributors

{{ range .Contributors }}
- {{ .Name }} ({{ .Role }})
{{ end }}
