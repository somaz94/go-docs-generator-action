# {{ .ProjectName }}

[![License](https://img.shields.io/github/license/{{ .GitHubRepo }})](https://github.com/{{ .GitHubRepo }})
![Latest Tag](https://img.shields.io/github/v/tag/{{ .GitHubRepo }})
![Top Language](https://img.shields.io/github/languages/top/{{ .GitHubRepo }}?color=green&logo={{ .Language }}&logoColor=b)
[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-{{ urlquery .MarketplaceName }}-blue?logo=github)](https://github.com/marketplace/actions/{{ .MarketplaceSlug }})

## Overview

{{ .Description }}

## Inputs

| Input | Required | Description | Default |
|-------|----------|-------------|---------|
{{ range .ConfigOptions }}
| `{{ .Name }}` | {{ .Required }} | {{ .Description }} | {{ .Default }} |
{{ end }}

## Example Workflow

Below is an example of how to use the **{{ .ProjectName }}** in your GitHub Actions workflow:

```yaml
{{ .UsageExample }}
```

## Features

{{ range .Features }}
- {{ . }}
{{ end }}

## Notes

{{ .Requirements }}

## License

This project is licensed under the {{ .License }} License - see the [LICENSE](LICENSE) file for details.

## Contributors

{{ range .Contributors }}
- {{ . }}
{{ end }}

## Contributing

{{ .ContributingMessage }}
