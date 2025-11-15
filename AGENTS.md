# Repository Guidelines

## Tasks and Sources

- README authoring/refactor:
  - Read: ./docs/agents/readme-template.md

## Project Structure & Module Organization
Everything lives under `cli/`. `main.go` bootstraps `cmd/`, which drives the Charmbracelet forms and orchestrates supporting packages: `config/` (persisted settings), `constants/` (template metadata), `operations/` (filesystem + Git tasks), `fileops/` (IO helpers), `ui/` (forms/styles), `validators/` (input guards), and `types/` (shared DTOs). Release packaging follows `.goreleaser.yaml`; avoid ad-hoc edits.

## Build, Test, and Development Commands
Work from `cli/`. `go run .` exercises the CLI end-to-end. `go build -o bin/odyssey ./...` compiles every package for manual smoke tests. `go test ./...` must pass before opening a PR; add `-run TestName` for focused loops. `go vet ./...` plus `gofmt -w $(rg --files -g'*.go')` keep static checks quiet. Only use `goreleaser build --snapshot` when verifying release metadata locally.

## Coding Style & Naming Conventions
Target Go 1.25 features only and rely on `gofmt` for layout—no manual alignment. Exported identifiers use PascalCase with short doc comments when packages are consumed externally (for example, `operations.Task`). Keep packages cohesive: new workflows belong in `operations/`, new form components in `ui/`. JSON/YAML keys stay lowercase-with-hyphens to match template metadata. Error messages begin lowercase and omit trailing punctuation.

## Testing Guidelines
Table-driven tests (`cases := []struct{...}{}`) cover config validation, template resolution, and task orchestration. Keep tests beside their code (`*_test.go`) and prefer in-memory helpers; use `t.TempDir()` when `operations` needs disk access. Target ≥80% coverage for any new package and add regression tests alongside bug fixes. Integration runs that hit `cmd.Execute` should stub filesystem and Git effects via `operations.Task`.

## Commit & Pull Request Guidelines
Commits follow Conventional Commit verbs (`feat:`, `fix:`, `chore:`) with imperative descriptions, mirroring recent history (`feat: refactor naming conventions…`). Keep each commit scoped to one change. Every PR should include intent, `go test ./...` evidence, screenshots or GIFs for UI work, and links to linked issues. Call out template or workflow changes in both the PR body and a `BREAKING CHANGE:` footer if user impact exists.

## Security & Configuration Tips
Never commit files under `~/.odyssey` or environment-specific credentials. When reproducing bugs, redact tenant names, AWS account numbers, and Terraform backend URIs in logs. Use throwaway access tokens while testing `operations.CloneProjectFiles`, and inspect generated projects before pushing to remote repositories.
