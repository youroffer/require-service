deep:
	dep-tree entropy cmd/main.go
.PHONY: deep

run:
	docker compose -f compose.yml up --build --no-log-prefix --attach require
.PHONY: run

cover-html: ### run test with coverage and open html report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
.PHONY: coverage-html

cover: ### run test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out
.PHONY: coverage