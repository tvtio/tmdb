build:
	go build

install:
	go install

test:
	go test -v -race  ./...

cover:
	go test -v -race  ./... -coverprofile=coverage.out
	goveralls -v -coverprofile=coverage.out -service=circle-ci -repotoken=xx
	rm coverage.out

lint:
	go vet ./...
	golint ./...

clean:
	go clean

deps: dev-deps
	# Dev dependencies

dev-deps:
	go get github.com/golang/lint/golint
	go get github.com/jstemmer/gotags
	go get github.com/axw/gocov/gocov
	go get github.com/mattn/goveralls

tags:
	gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .

dist-clean: clean
	rm -rf pkg src bin
