# Examples

This page demonstrates common Odyssey workflows end to end.

## Bootstrap a Serverless Project
```bash
odyssey-cli
# 1) Select the "serverless" template.
# 2) Pick the GitHub Actions workflow.
# 3) Enter variables (project code, AWS region, domain prefix, etc.).
# 4) Confirm the summary; Odyssey scaffolds infra + source + CI in seconds.
```

## Update Variables and Re-run
```bash
# Edit odyssey/project.yaml (e.g., change aws_region or directory), then run:
odyssey-cli
```

## Inspect Generated Outputs
```bash
odyssey-cli
# After the run completes, inspect:
# - odyssey/project.yaml for persisted answers
# - infra/ and project-files/ for Terraform + app code
# - .github/workflows/ for the generated CI/CD pipelines
```
