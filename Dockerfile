# syntax = docker/dockerfile:1

FROM golang:1.24.3-alpine AS build
RUN apk --no-cache add ca-certificates
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build

FROM scratch
COPY --from=build /app/wheres_my_bus /wheres_my_bus
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/wheres_my_bus"]
