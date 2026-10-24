# Security Policy

## Supported versions

Only the latest release receives security fixes.

## Reporting a vulnerability

Please do not open a public issue. Report privately through GitHub:
**Security → Report a vulnerability** on this repository.

Include the affected tool and version (`<tool> -version`), what you observed, and steps to reproduce. You can expect an acknowledgement within a few days. Fixes are released as a new version, with credit to the reporter unless you prefer otherwise.

## Scope

labdog tools such as `pi-watchdog` run with elevated privileges (they can shut a host down). Issues that allow unprivileged users or remote hosts to trigger that behaviour, or to execute commands, are in scope.
