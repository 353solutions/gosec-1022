all:
	$(error please pick a target)


install-sec-tools:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.14.0


test:
	govulncheck .
	gosec ./...
	go test -v ./...
