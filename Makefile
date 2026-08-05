run:
	export APPENV=local; go run -race cmd/main.go

update_dependency:
	go mod tidy

test:
	go clean -testcache && go test -race -v ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

mock:
	go generate ./internal/repositories/adaptor/... ./internal/services/...

compose-up:
	docker compose -f _compose/compose.yaml up -d

compose-down:
	docker compose -f _compose/compose.yaml down
