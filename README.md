[![Go Report Card](https://goreportcard.com/badge/github.com/zacwhy/go-diff-transactions)](https://goreportcard.com/report/github.com/zacwhy/go-diff-transactions)

# go diff transactions

## Getting started

```sh
$ go build

$ ./go-diff-transactions
Usage: ./go-diff-transactions FILE1 FILE2

$ ./go-diff-transactions diff/testdata/po.csv diff/testdata/local.csv
po
< [2020-01-02 100][0] [02 Jan 2020 po only S$1.00]
local
> [2020-01-03 100][0] [2020-01-03 100 local only]

# or

$ go run main.go diff/testdata/po.csv diff/testdata/local.csv
po
< [2020-01-02 100][0] [02 Jan 2020 po only S$1.00]
local
> [2020-01-03 100][0] [2020-01-03 100 local only]
```

## Development

Install `git gofmt pre-commit hook` to make sure Go files are formatted with gofmt before commits.
```sh
$ git config core.hooksPath hooks
```
