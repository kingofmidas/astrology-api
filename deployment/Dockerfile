FROM golang:1.17-alpine AS base

RUN apk add make

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make build-binaries


# run api app
FROM alpine:3 AS api

WORKDIR /root
COPY --from=base /src/bin/api ./bin/api

CMD ["./bin/api"]


# run collector app
FROM alpine:3 AS collector

WORKDIR /root
COPY --from=base /src/bin/collector ./bin/collector

CMD ["./bin/collector"]


# run migrate app
FROM alpine:3 AS migrate

WORKDIR /root

COPY --from=base /src/bin/migrate ./bin/migrate
COPY --from=base /src/migrations/ ./migrations/
COPY --from=base /src/scripts/ ./scripts/

RUN chmod +x ./scripts/*
RUN apk update && apk add --no-cache bash postgresql-client

CMD ["./scripts/wait-for-postgres.sh", "postgres", "./bin/migrate"]

