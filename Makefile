# dryLang build & dev tools
.PHONY: build test race vet lint clean release

VERSION ?= 1.0.0
LDFLAGS := -X drylang/cli.Version=$(VERSION)

build:
	go build -ldflags "$(LDFLAGS)" -o dry .

test:
	go test ./tools/tests/...

race:
	go build -race -o dry-race .
	go test -race ./tools/tests/...

vet:
	go vet ./...

lint: vet
	gofmt -l . | grep -v '^web/' | grep -v node_modules || true

clean:
	rm -f dry dry-race dry.exe dry-race.exe

release:
	go build -ldflags "$(LDFLAGS)" -o dry .
