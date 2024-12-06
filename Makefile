
init :
	git config core.hooksPath .githooks
	# https://github.com/golangci/golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

.PHONY : init
