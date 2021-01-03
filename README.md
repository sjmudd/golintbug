# golint bug

The bug seems to be that `golint` does not recognise a an iota-typed
constant comment on the same line as the definition.

The code builds happily yet `golint` complains:

```
$ golint main*.go
main.go:10:2: exported const GoLintBugConst1 should have comment (or a comment on this block) or be unexported
```

The comment is correctly provided on the same line as the constant definition.

- `go` reports `go version go1.15.5 darwin/amd64`
- `golint` installed by `go get -u golang.org/x/lint/golint` so up to date as of 2021-01-03.

See: https://github.com/golang/lint/issues/502
