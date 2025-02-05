# {{ .ProjectName }}

[![License](https://img.shields.io/github/license/{{ .GitHubRepo }})](https://github.com/{{ .GitHubRepo }})
![Latest Tag](https://img.shields.io/github/v/tag/{{ .GitHubRepo }})
![Top Language](https://img.shields.io/github/languages/top/{{ .GitHubRepo }}?color=green&logo={{ .Language }}&logoColor=b)
[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-{{ .MarketplaceName }}-blue?logo=github)](https://github.com/marketplace/actions/{{ .MarketplaceSlug }})

{{ .Description }}

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

## ⚙️ Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
{{ range .ConfigOptions }}
| {{ .Name }} | {{ .Type }} | {{ .Default }} | {{ .Description }} |
{{ end }}

## 📝 License

This project is licensed under the {{ .License }} License - see the [LICENSE](LICENSE) file for details.

## 👥 Contributors

{{ range .Contributors }}
- {{ .Name }} ({{ .Role }})
{{ end }}

## Contributing

{{ .Contributing }}
