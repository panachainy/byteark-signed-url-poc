dev:
	air

i:
	go get ./cmd

tidy:
	go mod tidy -v

t: test
test:
	go test ./...

tc: test.cov
test.cov:
	$(ENV_LOCAL_TEST) \
	go test -v -race -covermode=atomic -coverprofile=covprofile ./...

cov: cov.html
cov.htm:
	go tool cover -html=covprofile

cov.func:
	go tool cover -func=covprofile

f: fmt
fmt:
	go fmt ./...

w: wire
wire:
	wire ./...

g: generate
generate:
	go generate ./...

b: build
build:
	go build -o apiserver ./cmd
