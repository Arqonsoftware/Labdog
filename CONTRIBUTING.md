# Contributing to Labdog

Thanks for helping out. Labdog is a collection of small, dependency-free Go tools for home labs, so contributions that keep things simple are the most welcome.

## Ground rules

- **One command, one job.** Each command does a single thing well.
- **Static and portable.** Tools must build with `CGO_ENABLED=0` for linux amd64, arm64 and armv7, and should prefer the standard library over new dependencies.
- **Safe by default.** Anything destructive (shutdown, delete, reboot) needs a `-dry-run` flag and conservative defaults.
- **No hardcoded environments.** Addresses, MACs, paths and credentials come from flags or config, never from source.

## Getting started

You need Go 1.24 or newer.

    git clone https://github.com/Arqonsoftware/labdog.git
    cd labdog
    go vet ./...
    go test ./...
    go build ./cmd/labdog

## Making a change

1. Open an issue first for anything bigger than a small fix, so we can agree on the approach.
2. Branch from `main` and keep the change focused.
3. Run `gofmt`, `go vet ./...` and `go test ./...`. CI runs the same checks plus an armv7 cross-build.
4. Add or update tests and the README when behaviour changes.
5. Open a pull request using the template. Squash-merge is used, so the PR title becomes the commit message. Write it in the imperative ("Add retry to the watchdog command").

## Adding a new command

Labdog is one binary with a subcommand per tool.

1. Add `cmd/labdog/<name>.go` with a `run<Name>(args []string) int` function that owns its own `flag.FlagSet` (see `watchdog.go`).
2. Register it in the `switch` and the usage text in `cmd/labdog/main.go`.
3. Add a row and a flag table to the Commands section of `README.md`.

## Releases

Maintainers release by pushing a semver tag (`v1.2.3`). The Release workflow builds the binary with GoReleaser and publishes archives and checksums.

## Conduct and security

Participation is covered by the [Code of Conduct](CODE_OF_CONDUCT.md). Report vulnerabilities privately as described in [SECURITY.md](SECURITY.md), not in public issues.

Labdog is maintained by [Arqonsoftware](https://arqonsoftware.com).
