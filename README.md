# vfsgen-sample

An sample of using embedded file in Go.

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

## Reference
* https://golang.org/pkg/go/build/
* https://blog.golang.org/generate