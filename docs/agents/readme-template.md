# AGENTS.md — README Policy (Generic)

This repository enforces a standard README.md structure for any project or language. When creating or refactoring README.md, follow these rules exactly. If the existing README deviates, rewrite it to comply while preserving accurate project facts.

## Goals
- Consistent structure that’s easy to scan.
- Clear, runnable examples for any language.
- A concise, persuasive overview with light emphasis and relatable framing.

## Scope
- Applies to: /README.md
- Audience: All contributors and automation that writes/updates README.md
- Projects: Any language, framework, or stack

---

## Required Structure (in order)

1) Title
- H1: `# {Project Name}`
- Optional single‑line tagline below the H1.

2) Badges (optional but recommended)
- One badge per line, placed directly under the H1/tagline.
- Use only meaningful badges (build/test, package registry, license, docs).
- If none exist, omit this section entirely.

3) Overview
- H2: `## Overview`
- 2–4 concise sentences: what it is, the problem it solves, key outcomes.
- Use selective emphasis: bold phrases per sentence, never more than three consecutive bolded words.

4) Features
- H2: `## Features`
- Start with a “bridge” line in this section:
  - You may wonder… why is {Project} created when there are {alternatives}? Here's a detailed [comparison](./docs/COMPARISON.md) documenting the reason why I created this framework!
  - Create ./docs/COMPARISON.md with a more detailed comparison with alternatives
- Provide a flat bullet list (no nesting) of 4–6 positioning points after the bridge, each one line, each beginning with a bold keyword followed by a short explanation.

5) Installation / Setup
- H2: `## Installation`
- If language/tooling varies, provide the most common path and note alternatives.
- Include a single fenced block (`bash`, `powershell`, or language‑specific) with commands.
- If this is a library, show the dependency add instruction; if an app, show install/run prerequisites.

6) Quick Start
- H2: `## Quick Start`
- Short lead‑in sentence (“Get started in minutes”).
- Provide 1–2 minimal end‑to‑end examples:
  - Library: show import/usage snippet in the primary language.
  - App/Service/CLI: show startup + a basic invocation or API call.
- Use language‑tagged fences (e.g., `bash`, `sh`, `powershell`, `python`, `go`, `js`, `ts`, `java`).
- Keep examples runnable or trivially adaptable; avoid pseudo‑syntax.

7) How It Works
- H2: `## How It Works`
- 2–6 numbered steps that explain the main flow or architecture at a high level.
- Keep each step to one line if possible.

8) Configuration
- H2: `## Configuration`
- Document environment variables, config files, or flags.
- Show a small example with defaults in a fenced block (`env`, `yaml`, `json`, etc.).

9) Usage
- H2: `## Usage`
- Provide common commands/APIs with short descriptions.
- Include at least one example per primary interface (CLI/API/SDK), fenced with a language tag.

10) Project Structure (optional but recommended)
- H2: `## Project Structure`
- Short explanation plus a flat list of key directories/files and their purpose.
- Keep lines brief (one sentence max).

11) Performance or Limitations (optional)
- H2: `## Performance` or `## Limitations`
- Bullets describing notable performance traits, constraints, or trade‑offs.

12) Error Handling / Troubleshooting (optional)
- H2: `## Troubleshooting`
- Bullets for frequent issues and fixes.

13) Testing (optional)
- H2: `## Testing`
- Explain how to run tests with a fenced block of commands.
- Mention coverage or CI if relevant.

14) Documentation
- H2: `## Documentation`
- Flat bullet list linking to in‑repo docs or external docs (relative links preferred).

15) Roadmap (optional)
- H2: `## Roadmap`
- Brief bullets for upcoming milestones or features.

16) Security (optional but recommended for services/libraries)
- H2: `## Security`
- How to report vulnerabilities and a link to a SECURITY.md if present.

17) Contributing
- H2: `## Contributing`
- One line pointing to CONTRIBUTING.md and any code style or PR rules.

18) License
- H2: `## License`
- One line naming the license and linking to the LICENSE file.

19) Support
- H2: `## Support`
- Where to file issues, ask questions, or discuss (issue tracker, discussions, forum).

---

## Formatting Rules

- Headings: Use the exact heading levels above. Do not add extra top‑level sections beyond those listed; place custom content under Optional Sections.
- Lists: Use flat bullet lists. Do not nest bullets. Keep each bullet to a single line when possible.
- Code: Use fenced code blocks with language tags (`bash`, `sh`, `powershell`, `json`, `yaml`, `toml`, `python`, `go`, `js`, `ts`, `java`, `rust`, etc.).
- Links: Prefer relative links for in‑repo files (e.g., `./docs/guide.md`). Use absolute URLs for external resources.
- Images: Use relative paths for repo images (e.g., `./docs/diagram.png`). Provide a short caption line above each image if used.
- Tone: Concise and factual. Avoid marketing language. Emphasize concrete capabilities and instructions.

---

## Merge and Rewrite Policy

- If required sections are missing or out of order, rewrite the README to match this policy.
- Preserve valid technical details and working links; update structure and phrasing to comply.
- If a referenced file does not exist, keep the section but remove the broken link and add a TODO in the PR description listing missing files.

---

## Optional Sections

You may include these if genuinely useful to users:
- `## Examples` (extended examples or recipes)
- `## Architecture` (diagram and short explanation)
- `## Deployment` (instructions for container/K8s/Serverless)
- `## Compatibility` (language/runtime versions, OS support)
- `## Internationalization` (if applicable)

Place optional sections after “Usage” and before “Testing,” or near the most relevant required section.

---

## Self‑Test Checklist (must verify before commit)

- Title present and meaningful; badges (if any) immediately under H1.
- All required sections present, in the specified order, with correct heading levels.
- Code blocks fenced and language‑tagged; examples are runnable or trivially adaptable.
- No nested bullet lists; no orphan headings; links resolve for in‑repo docs.
- Optional sections only included if populated with useful content.
- Final pass with a markdown linter if available; fix headings, lists, and code fences.

---
