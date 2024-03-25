.PHONY: run test coverage coverage-html mock

run:
	go run main.go

run-all:
	docker compose up -d

mock:
	mockgen -source=nat/nat.go -destination=nat/mock_nat.go -package=nat

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
	open cover.html