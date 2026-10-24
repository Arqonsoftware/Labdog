# labdog 🐕

Small Go binaries for Arqonsoftware's servers and Raspberry Pis. Each tool lives in `cmd/<name>` and is released as a static binary for linux amd64, arm64 and armv7.

## Tools

| Tool | Purpose |
|---|---|
| `pi-watchdog` | Gracefully shuts the host down when a peer sharing its power (e.g. a Pi with no UPS) stops answering pings. |

## Build

    go build ./cmd/<name>

## Release

Push a tag; the Release workflow builds every tool with GoReleaser and attaches the archives and `checksums.txt` to a GitHub release.

    git tag v0.1.0 && git push origin v0.1.0

## Install

    gh release download --repo Arqonsoftware/labdog --pattern '*linux_amd64.tar.gz'
