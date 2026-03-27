# SpaceScale CLI

The command line interface for SpaceScale.

Note: everything in this repository is still stubbed and under active setup, including the CLI commands, release flow, and packaging automation.

## Development

```bash
make test
make build
make run
make release-snapshot
```

## Usage

Run the root stub:

```bash
go run .
```

Show version information:

```bash
go run . version
```

Check the login stub:

```bash
go run . login
```

Override the default environment:

```bash
SPACESCALE_ENV=staging go run .
```

## Releases

Create a local snapshot build with GoReleaser:

```bash
make release-snapshot
```

Publish a GitHub Release by pushing a version tag:

```bash
git tag v0.1.0
git push origin v0.1.0
```

The release workflow builds archives for macOS, Linux, and Windows and uploads them to GitHub Releases.

Homebrew tap publishing is also stubbed through GoReleaser for `spacescale/homebrew-tap`. Before using tagged releases for that path, add a `HOMEBREW_TAP_GITHUB_TOKEN` repository secret with permission to push to the tap repository.
