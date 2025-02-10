# Go URL Shortener

[![Ozon Tech](https://img.shields.io/badge/Ozon%20Tech-FFF.svg?&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDgiIGhlaWdodD0iNDgiIHZpZXdCb3g9IjAgMCA0OCA0OCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTE5Ljg4MDQgMTAuNzQ2MUM5LjAzODk5IDEzLjI0OTcgMi4xNjQxNSAyMS4yMTg4IDQuNDMyNTIgMjguNTE0NkM2LjcwMDg5IDM1LjgxMDUgMTcuMjc4MSAzOS43NDk1IDI4LjExOTYgMzcuMjQ2QzM4Ljk2MSAzNC43NDI0IDQ1LjgzNTkgMjYuNzczMyA0My41Njc1IDE5LjQ3NzVDNDEuMjk5MSAxMi4xODE2IDMwLjYzMzcgOC4yNjI5NSAxOS44ODA0IDEwLjc0NjFaTTI2LjIxNjIgMzEuMTI0MkMxOC43MjQyIDMyLjg1NDMgMTEuOTk3OSAzMC4zNjI4IDEwLjk4MTEgMjcuMDkyM0M5Ljk2NDIzIDIzLjgyMTcgMTQuMzQzOSAxOC43NjU3IDIxLjgzNTkgMTcuMDM1NkMyOS4zMjggMTUuMzA1NiAzNi4wNTQyIDE3Ljc5NyAzNy4wNzExIDIxLjA2NzVDMzguMDg3OSAyNC4zMzgxIDMzLjYyMDEgMjkuNDE0NCAyNi4yMTYyIDMxLjEyNDJaIiBmaWxsPSIjMDAxQTM0Ii8+Cjwvc3ZnPgo=)](https://ozon.tech/)
[![CI pipeline](https://github.com/anijackich/go-url-shortener/actions/workflows/ci.yml/badge.svg)](https://github.com/anijackich/go-url-shortener/actions/workflows/ci.yml)
[![Go version](https://img.shields.io/github/go-mod/go-version/anijackich/go-url-shortener)](https://github.com/anijackich/go-url-shortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/anijackich/go-url-shortener)](https://goreportcard.com/report/github.com/anijackich/go-url-shortener)

Simple Go API application for URL shortener service.

## ✨ Features

- [x] Convert long URL to a short link
- [x] Expand long URL from a short link
- [x] In-memory storage
- [x] PostgreSQL storage

## ⬇️ Clone & Setup

1. Clone this repository

   ```shell
   git clone https://github.com/anijackich/go-url-shortener
   ```

2. Go to `go-url-shortener` directory

   ```shell
   cd go-url-shortener
   ```

3. Rename `.env.example` to `.env`

   ```shell
   mv .env.example .env
   ```

4. Set environment variables

   | Variable             | Description                                         |
               |----------------------|-----------------------------------------------------|
   | `HOST`               | Host serving the API                                |
   | `PORT`               | Port serving the API                                |
   | `DOMAIN`             | Base domain for shortened links                     |
   | `LINK_CODE_LENGTH`   | Length of code in path of shortened links           |
   | `LINK_CODE_ALPHABET` | Acceptable chars of code in path of shortened links |
   | `POSTGRES_USER`      | PostgreSQL User                                     |
   | `POSTGRES_PASSWORD`  | PostgreSQL Password                                 |
   | `POSTGRES_HOST`      | PostgreSQL Host                                     |
   | `POSTGRES_PORT`      | PostgreSQL Port                                     |
   | `POSTGRES_DATABASE`  | PostgreSQL Database name                            |

## 💾 Storage

The application supports two possible options for storing links:

- `memory` temporary storage in memory
- `postgres` persistent storage in PostgreSQL

Once selected, replace `<STORAGE>` in following commands with one of these values.

## 🐋 Docker

### Run

```shell
docker run -it -p 8080:8080 --env-file .env anijack/go-url-shortener --storage <STORAGE>
```

### Run using Docker Compose (with PostgreSQL)

```shell
STORAGE=<STORAGE> docker compose up
```

However, it makes no sense to run application with `memory` storage type using Docker Compose, since PostgreSQL is
launched with it, which is necessary only for `postgres` storage type.

### Build from sources

```shell
docker build -t anijack/go-url-shortener .
```

## 🔨 Build

```shell
make build
```

## 🚀 Run

```shell
make run STORAGE=<STORAGE>
```

## ✔️ Code quality

### Run Formatter

```shell
make fmt
```

### Run Linter

```shell
make lint
```

[golangci-lint](https://github.com/golangci/golangci-lint) is required

### Run Tests

```shell
make test
```

Coverage report in `coverage.out`

