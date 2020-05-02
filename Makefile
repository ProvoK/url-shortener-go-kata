.PHONY: web
web:
	go build -o web ./cmd/web

.PHONY: run
run: web
	./web
