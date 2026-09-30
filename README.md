# ValidGen samples

These programs use [ValidGen](https://github.com/opencodeco/validgen) from a module of their own.

ValidGen reads `valid` and `validate` tags and writes `validator__.go` beside the structs. That file imports `github.com/opencodeco/validgen/types`. This module depends on `github.com/opencodeco/validgen` for that package.

ValidGen is experimental. These samples are too.

## Requirements

Go 1.24 or newer.

## Same package

`signup` defines `User` in the package that ValidGen generates into. `Note` has no validation tag, so there is no `NoteValidate`.

The `signup` generate line passes `-unmarshal-json`, so `User` also gets `UnmarshalJSON`. Decoding runs validation and returns `errors.Join` when the value is invalid.

```bash
go test ./signup
```

## Another package

`account` defines `Profile`. `cmd/account` imports `github.com/opencodeco/validgen-samples/account` and calls `ProfileValidate`.

The generated validator stays in `account`. The command does not import ValidGen internal packages.

```bash
go test ./account
go run ./cmd/account
```

## Run the CLI

Generate the two packages:

```bash
go run github.com/opencodeco/validgen@latest -unmarshal-json ./signup
go run github.com/opencodeco/validgen@latest ./account
```

Install the CLI and run the `go:generate` lines:

```bash
go install github.com/opencodeco/validgen@latest
go generate ./...
```

`go:generate` runs in the package directory. The argument `.` is that directory.

ValidGen overwrites `validator__.go` on every run. The copies in this repository are committed, so `go test ./...` does not need the CLI. After you change a struct tag, run the CLI again and commit the new file.

There is no ValidGen release tag yet. `@latest` is the `main` branch.
