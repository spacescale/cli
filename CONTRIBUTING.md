# Contributing to SpaceScale CLI

SpaceScale CLI is the command-line interface for SpaceScale. The CLI is still early, so contributions should stay small, reviewed, and scoped.

## Contribution Rules

- Open an issue or comment on an existing issue before starting larger work.
- Keep pull requests focused on one change.
- Do not add new dependencies without maintainer approval.
- Do not include secrets, private endpoints, test credentials, or internal SpaceScale infrastructure details.
- Keep public code and docs aligned with the current README.
- Make sure `make test` and `make build` pass before opening a pull request.

## Development

```bash
make test
make build
make run
```

Run the CLI locally:

```bash
go run .
go run . version
go run . login
```

## Pull Request Checklist

Before opening a pull request:

- [ ] The change is scoped to the issue.
- [ ] `make test` passes.
- [ ] `make build` passes.
- [ ] Public docs are updated if behavior changed.
- [ ] No secrets or private SpaceScale details are included.

## Contribution License

By contributing to this repository, you agree that your contributions are submitted under the repository license.

You also confirm that:

- you have the right to submit the contribution;
- the contribution is your own work or is otherwise legally allowed to be submitted;
- the contribution does not knowingly violate another person or company's intellectual property rights.
