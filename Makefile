.PHONY: web
web:
	go build -o web ./cmd/web

.PHONY: run
run: web
	./web

.PHONY: test
test:
	go test --coverprofile=coverage.out -v ./...
	go tool cover -func=coverage.out
