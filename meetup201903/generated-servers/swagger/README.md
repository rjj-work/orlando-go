# Ex-strs server generated with `swagger`

... TODO: say something here ...

## Generate the server

Input is `ex-strs.2.0.yaml` file


```bash
$ swagger generate server --target . --name ex-strs --spec ex-strs.2.0.yaml
2019/03/14 09:09:59 validating spec /Users/rjj/dev/go/src/github.com/rjj-work/orlando-go/meetup201903/generated-servers/swagger/ex-strs/ex-strs.2.0.yaml
2019/03/14 09:10:00 preprocessing spec with option:  minimal flattening
2019/03/14 09:10:00 building a plan for generation
2019/03/14 09:10:00 planning definitions
2019/03/14 09:10:00 planning operations
2019/03/14 09:10:00 grouping operations into packages
2019/03/14 09:10:00 planning meta data and facades
2019/03/14 09:10:00 rendering 0 models
2019/03/14 09:10:00 rendering 1 operation groups (tags)
2019/03/14 09:10:00 rendering 1 operations for example
2019/03/14 09:10:00 rendering 4 templates for operation ExStrs
2019/03/14 09:10:00 name field GetConcat
2019/03/14 09:10:00 package field example
2019/03/14 09:10:00 creating generated file "get_concat_parameters.go" in "restapi/operations/example" as parameters
2019/03/14 09:10:00 executed template asset:serverParameter
2019/03/14 09:10:00 name field GetConcat
2019/03/14 09:10:00 package field example
2019/03/14 09:10:00 creating generated file "get_concat_urlbuilder.go" in "restapi/operations/example" as urlbuilder
2019/03/14 09:10:00 executed template asset:serverUrlbuilder
2019/03/14 09:10:00 name field GetConcat
2019/03/14 09:10:00 package field example
2019/03/14 09:10:00 creating generated file "get_concat_responses.go" in "restapi/operations/example" as responses
2019/03/14 09:10:00 executed template asset:serverResponses
2019/03/14 09:10:00 name field GetConcat
2019/03/14 09:10:00 package field example
2019/03/14 09:10:00 creating generated file "get_concat.go" in "restapi/operations/example" as handler
2019/03/14 09:10:00 executed template asset:serverOperation
2019/03/14 09:10:00 rendering 0 templates for operation group ExStrs
2019/03/14 09:10:00 rendering support
2019/03/14 09:10:00 rendering 6 templates for application ExStrs
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 creating generated file "main.go" in "cmd/ex-strs-server" as main
2019/03/14 09:10:00 executed template asset:serverMain
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 creating generated file "embedded_spec.go" in "restapi" as embedded_spec
2019/03/14 09:10:00 executed template asset:swaggerJsonEmbed
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 creating generated file "server.go" in "restapi" as server
2019/03/14 09:10:00 executed template asset:serverServer
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 creating generated file "ex_strs_api.go" in "restapi/operations" as builder
2019/03/14 09:10:00 executed template asset:serverBuilder
2019/03/14 09:10:00 name field ExStrs
2019/03/14 09:10:00 package field operations
2019/03/14 09:10:00 creating generated file "doc.go" in "restapi" as doc
2019/03/14 09:10:00 executed template asset:serverDoc
2019/03/14 09:10:00 Generation completed!

For this generation to compile you need to have some packages in your GOPATH:

	* github.com/go-openapi/runtime
	* github.com/jessevdk/go-flags

You can get these now with: go get -u -f ./...```

## File structure

```bash
$ tree -d
.
├── cmd
│   └── ex-strs-server
└── restapi
    └── operations
        └── strings

5 directories
```

### Run it

```bash
$ go run cmd/ex-strs-server/main.go
2019/03/14 09:10:10 Serving ex strs at http://127.0.0.1:56667
2019/03/14 09:10:10 the required flags `--tls-certificate` and `--tls-key` were not specified
exit status 1
gandalf15:ex-strs rjj$ go run cmd/ex-strs-server/main.go --scheme=http
2019/03/14 09:10:28 Serving ex strs at http://127.0.0.1:56671
```