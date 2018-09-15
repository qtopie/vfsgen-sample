all: build

setup:
	go get -u github.com/shurcooL/vfsgen

tools:
	go get -u golang.org/x/lint/golint

lint:
	golint

fmt: 
	go fmt ./... 

build:
	cd data && go generate -tags dev && cd -
	go build -ldflags '-X main.version=1.0'

install:
	go install .

clean:
	go clean
	rm -f data/assets_vfsdata.go
