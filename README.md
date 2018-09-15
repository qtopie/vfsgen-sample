# vfsgen-sample

An sample of using embedded files in Go.

## Build
```bash
# with makefile (recommended)
make

# with commands
cd data && go generate -tags dev && cd -
go build
```

## Cleanup
```bash
# makefile
make clean

# or
go clean
rm -f data/assets_vfsdata.go
```

## Versioning

Using `-ldflags`:

```bash
go build -ldflags '-X main.version=1.0'
```

## Run

Run `./vfsgen-sample` after building, then open http://localhost:7070 in your brower to check the served embedded files.

Also view the console log to check whether reading files from `vfsgen` is okay.

## Reference
* https://golang.org/pkg/go/build/
* https://blog.golang.org/generate