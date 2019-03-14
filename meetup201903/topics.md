# Something about Go (TBD)


## Go Gotcha
1. pointer vs. non-pointer receiver
1. nil receiver possible
1. return statement with vars not always necessary
   - Is this new in Go 1.11.x ?

## Closure Examples
1. Checker
1. Test cleanup: CUT

## Channels
1. Channel state
1. Concurrency
    1. Tony Hoare's CSP 1978: http://usingcsp.com/cspbook.pdf
    1. Google I/O 2012 - Go Concurrency Patterns: https://www.youtube.com/watch?v=f6kdp27TYZs

## Generated Servers

1. Cobra
    1. [spf13/cobra](https://github.com/spf13/cobra/blob/master/cobra)
    1. [spf13/viper](https://github.com/spf13/viper)
    1. code generation

        ```bash
        cobra init ex-strs -a 'rjj.work@gmail.com' -l MIT
        # ... mv used to put generated code where I want it ...
        cobra add concat
        ```

1.  Swagger
    1. [swagger.io](https://swagger.io)
    1. [swagger editor](https://editor.swagger.io/)
    1. [Homebrew to install on MacOS](https://brew.sh)
    1. code generation

        ```bash
        swagger generate server --target . --name ex-strs --spec ex-strs.2.0.yaml
        ```
