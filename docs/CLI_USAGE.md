# Odyssey CLI Usage

This guide explains how to drive the Odyssey CLI end-to-end, from collecting configuration to generating a scaffolded project, with concrete examples of the output you should expect on a healthy run.

## Prerequisites
- Git installed.
- Writable locations for `~/.odyssey/globals.json` (global preferences) and `<project>/odyssey/project.yaml` (per-project metadata).

## Configuration Flow – `odyssey-cli`
Run the CLI via `odyssey-cli` or execute the built binary. The wizard proceeds as follows:
1. **Welcome + health check** – surfaces missing config files or corruption warnings.
2. **Global settings** – collects Terraform backend info, AWS defaults, and any secrets that apply to every project.
3. **Project metadata** – prompts for project code, target directory, and template selection; the directory auto-fills from the code but can be overridden.
4. **Workflow & variables** – once a template/type is chosen the form unlocks workflow variants plus all required variables (region, domain name, etc.). Required fields cannot be skipped.
5. **Confirmation** – shows a formatted summary. Approving this step writes `~/.odyssey/globals.json` and `odyssey/project.yaml`.

### Example confirmation output
```text
$ odyssey-cli
╭───────────────────────────────────────────────╮
│ Odyssey Project Configuration                 │
│                                               │
│ Project Code:     sample-api                  │
│ Directory:        ./sample-api                │
│ Template:         Standard API (ECS)          │
│ Project Type:     REST API (scheduled)        │
│ Workflow:         Multi-environment           │
│ Environments:     preprod, production         │
╰───────────────────────────────────────────────╯
```
Your exact field names depend on the template’s variable catalog, but the border and ordering will match.


### Confirm Setup
```text
$ odyssey-cli
⠋ Creating project directory...
⠙ Cloning project files...
⠹ Initializing Git repository...
⠸ Adding project submodules...
✔ Project files ready

╭───────────────────────────────────────────────╮
│ Odyssey Project Configuration                 │
│                                               │
│ Project Code:     sample-api                  │
│ Directory:        ./sample-api                │
│ Template:         Standard API (ECS)          │
│ Project Type:     REST API (scheduled)        │
│ Workflow:         Multi-environment           │
│ Environments:     preprod, production         │
╰───────────────────────────────────────────────╯
```
Spinner glyphs change while work is in progress; the command exits non-zero if any task fails, surfacing the task description in the error message. Re-run `odyssey-cli` whenever you need to switch templates or workflows.