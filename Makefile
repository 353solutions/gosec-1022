gosec_ver=2.14.0
gosec_file=gosec_$(gosec_ver)_$(shell go env GOOS)_$(shell go env GOARCH).tar.gz
gosec_url=https://github.com/securego/gosec/releases/download/v$(gosec_ver)/$(gosec_file)

all:
	$(error please pick a target)


install-sec-tools:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	curl -L $(gosec_url) | tar xz -C $(shell go env GOPATH)/bin gosec


test:
	govulncheck .
	gosec ./...
	go test -v ./...
