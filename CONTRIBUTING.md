# Contributing to labdog

Thanks for helping out. labdog is a collection of small, dependency-free Go tools for home labs, so contributions that keep things simple are the most welcome.

## Ground rules

- **One tool, one job.** Each tool lives in `cmd/<name>` and does a single thing well.
- **Static and portable.** Tools must build with `CGO_ENABLED=0` for linux amd64, arm64 and armv7, and should prefer the standard library over new dependencies.
- **Safe by default.** Anything destructive (shutdown, delete, reboot) needs a `-dry-run` flag and conservative defaults.
- **No hardcoded environments.** Addresses, MACs, paths and credentials come from flags or config, never from source.

## Getting started

You need Go 1.24 or newer.

    git clone https://github.com/Arqonsoftware/labdog.git
    cd labdog
    go vet ./...
    go test ./...
    go build ./cmd/<tool>

## Making a change

1. Open an issue first for anything bigger than a small fix, so we can agree on the approach.
2. Branch from `main` and keep the change focused.
3. Run `gofmt`, `go vet ./...` and `go test ./...`. CI runs the same checks plus an armv7 cross-build.
4. Add or update tests and the README when behaviour changes.
5. Open a pull request using the template. Squash-merge is used, so the PR title becomes the commit message. Write it in the imperative ("Add retry to pi-watchdog").

## Adding a new tool

1. Create `cmd/<name>/main.go` with a `-version` flag that prints the `version` variable set at build time (see `cmd/pi-watchdog`).
2. Add a build entry to `.goreleaser.yaml`.
3. Add a row to the tools table in `README.md`, with usage and flags.

## Releases

Maintainers release by pushing a semver tag (`v1.2.3`). The Release workflow builds all tools with GoReleaser and publishes archives and checksums.

## Conduct and security

Participation is covered by the [Code of Conduct](CODE_OF_CONDUCT.md). Report vulnerabilities privately as described in [SECURITY.md](SECURITY.md), not in public issues.

labdog is maintained by [Arqonsoftware](https://arqonsoftware.com).
