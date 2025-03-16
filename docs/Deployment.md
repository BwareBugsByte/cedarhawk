# Deployment and CI/CD Integration Guide

This guide explains how to build, package, and deploy CedarHawk using the provided Makefile and CI/CD pipeline.

## Build Script

A Makefile is provided at the repository root. It supports the following targets:

- **make test**: Runs all unit tests.
- **make build**: Compiles CedarHawk for multiple platforms.
- **make release**: Packages the built binaries into archives for release.
- **make clean**: Cleans up all build and release artifacts.

### Example Usage

To run tests and build binaries:
```bash
make all
```
To package the release:
```bash
make release
```

The binaries will be output to the build/ directory, and the packaged artifacts will be available in the release/ directory.
## CI/CD Pipeline

A GitHub Actions workflow is configured in .github/workflows/ci.yml to:

  - Run tests on every commit or pull request.
  - Build the binaries for multiple platforms.
  - Package the binaries into archives.
  - Upload release artifacts when a new release is created.

## Workflow Triggers

  - **Pushes and Pull Requests:** The workflow runs on any push or pull request to the main branch.
  - **Releases:** When you create a new GitHub release, the pipeline automatically builds and packages the binaries, and uploads them as artifacts.

## Release Management

  1. Creating a Release:
      - Tag your release in Git:
        ```bash
        git tag -a v1.0.0 -m "Release version 1.0.0"
        git push origin v1.0.0
        ```

      - On GitHub, create a new release using the tag.
      - The CI/CD pipeline will build the binaries and attach them as release artifacts.

  2. Accessing Pre-Built Binaries:
      - Navigate to the repository's Releases page.
      - Download the desired binary archive for your platform.

## Troubleshooting

  - Build Failures:
    Check the CI/CD logs in the GitHub Actions tab for details.
  - Test Failures:
    Run go test ./... locally to diagnose issues.
  - Artifact Issues:
    Ensure the release/ directory is writable and that the Makefile is correctly configured.
