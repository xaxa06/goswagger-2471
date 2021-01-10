Reproducer for go swagger [issue 2471](https://github.com/go-swagger/go-swagger/issues/2471).

Deps:
- needs go swagger >= v0.25.0 in $PATH
- tested with go 1.15.6 (Linux/Debian)

```sh
>make

[...]
go build -o bin/bug
# github.com/swagger2471/pkg/example/client/operations
pkg/example/client/operations/example_post_parameters.go:129:2: joinedFoo declared but not used
make: *** [Makefile:4: all] Error 2
```





