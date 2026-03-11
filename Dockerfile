# syntax=docker/dockerfile:1
FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/bin/api /app/api
COPY --from=build /app/migrations /app/migrations

EXPOSE 8080
CMD ["/app/api"]
