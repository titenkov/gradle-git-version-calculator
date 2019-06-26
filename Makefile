build: test
	go build ggvc.go

test:
	go test .

package: build
	tar --exclude='./.git' --exclude='./Makefile' --exclude='./README.md' \
			 -zcvf "gradle-git-version-calculator-0.0.1.tar.gz" .