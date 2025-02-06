# Go Docs Generator Action

[![License](https://img.shields.io/github/license/somaz94/go-docs-generator-action)](https://github.com/somaz94/go-docs-generator-action)
![Latest Tag](https://img.shields.io/github/v/tag/somaz94/go-docs-generator-action)
![Top Language](https://img.shields.io/github/languages/top/somaz94/go-docs-generator-action?color=green&logo=go&logoColor=b)
[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Go%20Docs%20Generator%20Action-blue?logo=github)](https://github.com/marketplace/actions/go-docs-generator-action)

## Overview

The **Go Docs Generator Action** is a GitHub Action that automates documentation generation for your projects. Written in Go, it provides a flexible and efficient way to create and maintain documentation in your GitHub repositories. This action is particularly useful for generating consistent README files and other documentation with customizable templates.

## Inputs

| Input | Required | Description | Default |
|-------|----------|-------------|---------|
| `output_path` | No | Output documentation path | docs |
| `template_type` | No | Template type (readme, api, examples) | readme |
| `template_names` | No | Comma-separated list of template names | default |
| `format` | No | Output format | markdown |
| `project_name` | Yes | Project name | - |
| `github_repo` | Yes | GitHub repository name | - |
| `marketplace_name` | Yes | Marketplace name | - |
| `marketplace_slug` | Yes | Marketplace slug | - |
| `description` | No | Project description | - |
| `features` | No | Comma-separated list of features | - |
| `requirements` | No | Project requirements | - |
| `installation_steps` | No | Installation steps | - |
| `usage_example` | No | Usage example | - |
| `language` | No | Primary language | go |
| `license` | No | License type | MIT |
| `contributing_message` | No | Contributing message | Contributions are welcome! |
| `config_options` | No | Configuration options | - |
| `contributors` | No | List of contributors | - |
| `action_config` | No | Include action.yml inputs in config | false |

## Example Workflow

Below is an example of how to use the **Go Docs Generator Action** in your GitHub Actions workflow:

```yaml
uses: somaz94/go-docs-generator-action@v1
with:
  output_path: docs
  template_type: readme
  template_names: default,detailed
  format: markdown
  project_name: 'My Project'
  github_repo: 'username/my-project'
  marketplace_name: 'My Action'
  marketplace_slug: 'my-action'
  description: 'My awesome project'
  features: 'Feature 1,Feature 2,Feature 3'
  requirements: 'Go 1.21 or higher'
  installation_steps: 'go get github.com/somaz94/go-docs-generator-action'
  usage_example: |
    uses: somaz94/go-docs-generator-action@v1
    with:
      output_path: docs
      template_type: readme
      template_names: default,detailed
      format: markdown
      project_name: 'My Project'
      github_repo: 'username/my-project'
      marketplace_name: 'My Action'
      marketplace_slug: 'my-action'
...
```

## Features

- Written in Go for better performance and reliability
- Multiple template support
- Customizable output formats
- Flexible configuration options
- Automatic README generation
- Support for multiple documentation types
- GitHub Marketplace integration
- Detailed configuration options

## Notes

- Templates are located in the `templates` directory
- Multiple templates can be generated at once
- Supports both simple and detailed README formats
- Configuration options can be customized via inputs
- Action inputs can be included in configuration documentation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.