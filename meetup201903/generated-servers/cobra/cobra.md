# Cobra https://github.com/spf13/cobra

## Install

[spf13/cobra generate](https://github.com/spf13/cobra/blob/master/cobra/README.md)

```bash
gandalf15:cobra rjj$ export GO111MODULE=on
gandalf15:cobra rjj$ go get -u github.com/spf13/cobra
gandalf15:cobra rjj$ which cobra
/Users/rjj/dev/go/bin/cobra

gandalf15:cobra rjj$ cobra init ex-strs -a 'rjj.work@gmail.com' -l MIT
Using config file: /Users/rjj/.cobra.yaml
Your Cobra application is ready at
/Users/rjj/dev/go/src/ex-strs.

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.
```

Use of `GO111MODULE` either `on` or `auto` did not affect the placement of the output

- Moved code from its generated location:
-- `mv /Users/rjj/dev/go/src/ex-strs .`
-- Had to update import for new location

```
gandalf15:ex-strs rjj$ tree
.
├── LICENSE
├── cmd
│   └── root.go
└── main.go

1 directory, 3 files
```

## Create CLI with command `concat`

```
gandalf15:ex-strs rjj$ cobra add concat
Using config file: /Users/rjj/.cobra.yaml
concat created at /Users/rjj/dev/go/src/github.com/rjj-work/orlando-go/meetup201903/generated-servers/cobra/ex-strs/cmd/concat.go
gandalf15:ex-strs rjj$ tree
.
├── LICENSE
├── cmd
│   ├── concat.go
│   └── root.go
└── main.go

1 directory, 4 files
```

