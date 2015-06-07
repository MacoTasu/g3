# g3

## Description

Go + S3 = g3. g3 is CLI for dealing with S3.

Now it has been implemented only get command.
Also it will implement at any time other command.

## Usage
### get

```bash
$ g3 get <bucketname> <target file or directory>
```

If you want to download the 'test.txt' files from the 'test' directory, you can be achieved with the following command:

```bash
g3 get <bucketname> test/test.txt
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/MacoTasu/g3
$ go build
$ go install
$ aws configure
$ > AWS_ACCESS_KEY_ID
$ > AWS_SECRET_KEY
```

## Contribution

1. Fork ([https://github.com/MacoTasu/g3/fork](https://github.com/MacoTasu/g3/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[MacoTasu](https://github.com/MacoTasu)
