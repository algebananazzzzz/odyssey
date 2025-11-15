# Odyssey Framework vs Alternatives

| Alternative | Where It Shines | Why Odyssey Framework Differentiates |
|-------------|-----------------|-------------------------------------|
| AWS Amplify | Fast front-end centric scaffolds tied to AWS services | Odyssey stays cloud agnostic with Terraform modules and GitHub Actions that you fully own, so you can move workloads between providers without rewrites. |
| Serverless Framework | Rich plugin ecosystem for Lambda-focused stacks | Odyssey bundles CI/CD, IaC, and opinionated best practices (TDD + trunk-based) out of the box without depending on plugins or vendor-managed runtimes. |
| Terraform (vanilla) | Strong IaC engine with broad provider coverage | Odyssey layers a guided TUI, project templates, and variable injection to produce production-ready repos in seconds instead of hand-coding every module. |
| Pulumi | Polyglot IaC with familiar programming languages | Odyssey keeps everything in standard Terraform + GitHub Actions while still giving you customizable templates, so teams retain cloud portability without learning a new IaC language. |
| GitHub Actions alone | Flexible workflow automation and CI/CD primitives | Odyssey ships workflows plus infrastructure and application code together, letting teams scaffold an entire project—including pipelines—through one CLI run. |

## Why Odyssey Feels Different
- **One interactive flow:** `odyssey-cli` is the single entrypoint—no flags, no scripts—so every scaffold follows the same validated steps.
- **Terraform-first + GitHub Actions:** Infrastructure and CI/CD live side by side, keeping teams cloud agnostic while still opinionated about quality gates.
- **Built-in guardrails:** TDD defaults, trunk-based conventions, and persisted configs land in every repo, reducing onboarding time for new contributors.
- **Template ownership:** You can fork and evolve templates at will, layering in language choices (Go APIs, Puppeteer jobs, etc.) without waiting on plugin ecosystems.

Odyssey Framework combines Terraform portability, GitHub Actions pipelines, and accelerator templates so teams can scaffold production-grade projects with consistent guardrails in seconds.
