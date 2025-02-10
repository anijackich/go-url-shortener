# syntax=docker/dockerfile:1

ARG GO_VERSION=1.23.4
ARG ALPINE_VERSION=3.21
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY api/ api/
COPY cmd/ cmd/
COPY docs/ docs/
COPY internal/ internal/
COPY pkg/ pkg/

RUN --mount=type=cache,target=/go/pkg/mod  \
    go build -o go-url-shortener ./cmd/main.go

FROM alpine:${ALPINE_VERSION} as runtime

WORKDIR /app

ARG UID=10001
ARG GID=10001
RUN addgroup --gid ${GID} appgroup && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    --ingroup appgroup \
    appuser

USER appuser

COPY --from=build /app/go-url-shortener .

EXPOSE 8080

ENTRYPOINT ["./go-url-shortener"]