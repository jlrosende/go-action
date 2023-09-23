
.PHONY: cmd_deploy
cmd_deploy:
	go run main.go deploy

.PHONY: cmd_release
cmd_release:
	go run main.go deploy

.PHONY: cmd_test
cmd_test:
	go run main.go deploy

.PHONY: test
test:
	go test -v -cover -short ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: clean
clean:
	echo "Clean generated files"

.PHONY: docs
docs:
	echo "Gen docs"
