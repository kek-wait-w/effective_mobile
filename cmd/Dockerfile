
FROM golang:1.22.1-alpine AS build-stage


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /app/main /

ENTRYPOINT ["/main"]
