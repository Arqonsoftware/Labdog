<p align="center">
  <img src="assets/logo.svg" alt="Labdog" width="560">
</p>

<p align="center">
  <a href="https://arqonsoftware.com">arqonsoftware.com</a>
</p>

<p align="center">
  <a href="https://github.com/Arqonsoftware/labdog/actions/workflows/ci.yml"><img alt="CI" src="https://github.com/Arqonsoftware/labdog/actions/workflows/ci.yml/badge.svg"></a>
  <a href="https://github.com/Arqonsoftware/labdog/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/Arqonsoftware/labdog?sort=semver"></a>
  <a href="LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
</p>

**Labdog** is a collection of small, single-purpose Go binaries that keep home-lab servers and Raspberry Pis well behaved. Each tool is a static binary with no runtime dependencies, released for Linux on amd64, arm64 and armv7.

## Commands

Labdog ships as a single binary, `labdog`, with a subcommand per tool.

| Command | What it does |
|---|---|
| [`labdog watchdog`](cmd/labdog/watchdog.go) | Shuts the host down gracefully when a peer that shares its power (for example a Pi with no UPS) stops answering pings. |
| `labdog version` | Prints the version. |

### labdog watchdog

A Pi on the same power circuit as your server vanishes the moment mains power drops. `labdog watchdog` treats that as an outage signal and shuts the server down cleanly. It only arms itself after it has seen the peer up, so a peer that is down at boot can't cause a shutdown loop.

    labdog watchdog -peer 192.168.0.2 -fail-after 60s

| Flag | Default | Description |
|---|---|---|
| `-peer` | required | Address of the power-loss sensor |
| `-fail-after` | `60s` | How long the peer must be unreachable before shutting down |
| `-interval` | `10s` | Ping interval |
| `-dry-run` | `false` | Log instead of shutting down |

It needs root (or permission to run `shutdown`) and a `ping` binary. To bring the server back after power returns, have the peer send a Wake-on-LAN packet at boot and enable "restore on AC power loss" in the server's BIOS.

Example systemd unit:

    [Unit]
    Description=Labdog watchdog
    After=network-online.target

    [Service]
    ExecStart=/usr/local/bin/labdog watchdog -peer 192.168.0.2
    Restart=always

    [Install]
    WantedBy=multi-user.target

## Install

Download a release archive for your architecture (`amd64`, `arm64` or `armv7`) and verify it against `checksums.txt`:

    gh release download --repo Arqonsoftware/labdog --pattern 'labdog_*_linux_amd64.tar.gz'
    tar xzf labdog_*_linux_amd64.tar.gz
    sudo install labdog /usr/local/bin/

Or build from source with Go 1.24+:

    go install github.com/Arqonsoftware/labdog/cmd/labdog@latest

## Development

    go vet ./...
    go build ./cmd/labdog

New commands are subcommands of `labdog`; see [CONTRIBUTING.md](CONTRIBUTING.md).

## Releasing

Push a tag and the Release workflow builds the binary with GoReleaser and publishes a GitHub release with archives and checksums:

    git tag v0.1.0 && git push origin v0.1.0

## Contributing

Issues and pull requests are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) first, and report vulnerabilities as described in [SECURITY.md](SECURITY.md).

## License

[MIT](LICENSE) © [Arqonsoftware](https://arqonsoftware.com)

Labdog is built and maintained by [Arqonsoftware](https://arqonsoftware.com).
