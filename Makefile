all: fix vet lint test

fix:
	go fix ./...

vet:
	go vet ./...

test: lint
	go test -v -cover ./...

lint:
	gometalinter ./... --vendor
