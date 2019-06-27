clean:
	rm -rf dist

build: clean test
	go build ggvc.go

test: clean
	go test .

release:
	goreleaser