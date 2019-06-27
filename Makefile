clean:
	rm -rf dist

build: clean test
	go build ggvc.go

test:
	go test .

release:
	goreleaser

package: build
	tar --exclude='./.git' --exclude='./Makefile' --exclude='./README.md' \
			 -zcvf "gradle-git-version-calculator-0.0.1.tar.gz" .