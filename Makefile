GOOS?=linux
GOARCH?=amd64

.SILENT:

.PHONY:
build-api:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/api ./cmd/api/main.go

.PHONY:
build-collector:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/collector ./cmd/collector/main.go

.PHONY:
build-migrate:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/migrate ./cmd/migrate/main.go

.PHONY:
build-binaries: build-api build-collector build-migrate

.PHONY:
run:
	docker-compose -f deployment/docker-compose.yml --project-directory . up

.PHONY:
build:
	docker-compose -f deployment/docker-compose.yml --project-directory . build

.PHONY:
down:
	docker-compose -f deployment/docker-compose.yml --project-directory . down
