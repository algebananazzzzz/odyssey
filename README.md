# Odyssey Framework

[![GitHub Release](https://img.shields.io/github/v/release/algebananazzzzz/odyssey?label=release)](https://github.com/algebananazzzzz/odyssey/releases)
[![License](https://img.shields.io/github/license/algebananazzzzz/odyssey)](LICENSE)

**Odyssey Framework** empowers developers to launch production-grade AWS projects complete with **CI/CD pipelines and Infrastructure-as-Code** in just seconds. Whether you're building a serverless Golang API or an automated Puppeteer job, Odyssey scaffolds an architecture that’s **ready to ship in seconds**. Its modular **Terraform** setup defines robust **AWS infrastructure**, while **GitHub Actions** provides fully customizable workflows without any vendor lock-in. All CI workflows follow modern DevOps practices inspired by **Test-Driven-Development (TDD)** and **Trunk-Based Development (TBD)**, ensuring you start right, right from the start.


## Features
You may wonder… why is Odyssey Framework created when there are tools like the Amplify, Serverless Framework? Here's a detailed [comparison](./docs/COMPARISON.md) documenting the reason why I created this framework!

- **Ship in Seconds:** Odyssey bundles CI/CD pipelines, Infrastructure-as-Code, testing workflows, and project scaffolding so you can generate a real-world, deploy-ready AWS project instantly.
- **Cloud-agnostic and vendor-neutral:** Powered by Terraform and GitHub Actions, Odyssey uses open, industry-standard tooling—giving you full control with no dependency on proprietary or locked-in workflows.
- **Best practices baked in:** Modern DevOps principles — TDD and TBD workflows are incorporated from day one so teams start with reliable processes instead of reinventing them.
- **10+ templates available:** Odyssey includes ready-to-launch templates for websites, REST APIs, event-driven functions, scheduled jobs, browser automation, and more—covering everyday production workloads right out of the box.

## Installation
Odyssey provides prebuilt CLI binaries for all major platforms. Download the archive that matches your system:

### **Supported Platforms**

| OS        | Architecture | Download |
|-----------|--------------|----------|
| **macOS** | arm64        | [Download](https://github.com/algebananazzzzz/odyssey/releases/latest/download/odyssey-cli_Darwin_arm64.tar.gz) |
| **macOS** | amd64        | [Download](https://github.com/algebananazzzzz/odyssey/releases/latest/download/odyssey-cli_Darwin_amd64.tar.gz) |
| **Linux** | amd64        | [Download](https://github.com/algebananazzzzz/odyssey/releases/latest/download/odyssey-cli_Linux_amd64.tar.gz) |
| **Linux** | arm64        | [Download](https://github.com/algebananazzzzz/odyssey/releases/latest/download/odyssey-cli_Linux_arm64.tar.gz) |
| **Windows** | amd64      | [Download](https://github.com/algebananazzzzz/odyssey/releases/latest/download/odyssey-cli_Windows_amd64.zip) |

---

### **Install (macOS / Linux)**

```bash
tar -xzf odyssey-cli_<os>_<arch>.tar.gz
chmod +x odyssey
sudo mv odyssey /usr/local/bin/odyssey
```

### **Install (Windows PowerShell)**

```powershell
Expand-Archive odyssey-cli_Windows_amd64.zip -DestinationPath .
Move-Item odyssey.exe "C:\Windows\System32\odyssey.exe"
```

### **Verify Installation**

```bash
odyssey --version
```

## Quick Start
Get started in minutes.

1. Launch the wizard:

```bash
odyssey
```

## How It Works

1. Resolve template metadata from the template GitHub repositories (`cli/constants/project.config.yaml`).
2. Retrieve all supported projects/templates defined by `odyssey.template-config.yaml` in the selected repository.
3. Render an interactive form (TUI) to collect project-specific variables with inline validation.
4. Execute project generation — create folders, fetch template files, inject variables, add submodules, initialize Git...
5. Persist config to `~/.odyssey/globals.json` and `odyssey/project.yaml`.

## Project Structure
- `cli/cmd` — CLI entrypoint and orchestrator.
- `cli/config` — loading and saving global/project configuration.
- `cli/constants` — template catalog definitions.
- `cli/operations` — filesystem and Git task implementations.
- `cli/ui` — Charmbracelet forms, prompts, and summary formatting.
- `docs/` — usage guides and comparisons.
- `AGENTS.md` — contributor standards.

## Performance
- Static binary (CGO disabled) with minimal startup overhead.
- Template metadata caching reduces repeated network calls.
- Spinner-backed task pipeline keeps progress clear during scaffolds.
- Works offline once templates are cached locally.

## Troubleshooting
- **Missing template repo:** Ensure the Git repo and ref in `project.config.yaml` are accessible from your environment.
- **Validation failures:** Update the values in `odyssey/project.yaml` or rerun the wizard to provide missing variables.
- **Permission errors:** Make sure the install path is writable or run the CLI from a directory you control.

## Testing
Run the full suite from the CLI workspace.

```bash
cd cli
go test ./...
```

## Documentation
- [CLI Usage Guide](./docs/CLI_USAGE.md)
- [Template Comparison](./docs/COMPARISON.md)
- [Examples](./docs/EXAMPLES.md)

## Roadmap
- Expand the template catalog (ECS jobs, VPC baselines, data pipelines).
- Publish architecture and Terraform module deep dives.
- Automate release builds end-to-end with GitHub Actions + GoReleaser.

## Contributing
Fork the repo, branch from `main`, and run `go test ./...` before submitting a PR. See [AGENTS.md](./AGENTS.md) for detailed expectations.

## License
Distributed under the [MIT License](./LICENSE).

## Support
Open an issue with repro steps, sanitized template descriptors, and CLI logs. For private templates, redact credentials or tenant identifiers before sharing.
