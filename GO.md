## Frameworks choice

[Gin](https://github.com/gin-gonic/gin) and [GORM](https://github.com/go-gorm/gorm) were chosen to implement the project
for their undemanding nature and relative simplicity for use in small projects

## Implementation solutions

### Mutual uniqueness of shortened and original url

#### Memory implementation

A bidirectional hashmap was used for the solution, as it allows checking for the existence of a link both from the
original url and from the shortened one code in the optimal time.

```go
type LinkRepository struct {
    linksByCode map[string]*models.Link
    codesByUrl  map[string]string
    ...
}
```

#### PostgreSQL implementation

In case of database implementation, two unique indexes are used for the links table (original url, shortened link code),
which allows to achieve the same effect.

```go
type Link struct {
    ...
    ShortCode string `gorm:"uniqueIndex;not null"`
    URL       string `gorm:"uniqueIndex;not null"`
}
```

### Short URLs generation

To generate short links, the base domain passed to .env is used, to which the generated code is added to the path.

The code is generated using the built-in `math/rand` library. For each element of an array of a given length, one of the
characters from alphabet is randomly selected, and then everything is casts to string.

```go
r := rand.New(rand.NewSource(time.Now().UnixNano()))

b := make([]byte, length)
for i := range b {
    b[i] = alphabet[r.Intn(len(alphabet))]
}
```

### Project structure

```text
go-url-shortener
├── api
│   └── swagger
│       └── docs.go
├── cmd
│   └── main.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal
│   ├── config
│   │   └── config.go
│   ├── handlers
│   │   └── links.go
│   ├── models
│   │   └── links.go
│   ├── repository
│   │   ├── errors.go
│   │   ├── memory
│   │   │   └── memory.go
│   │   ├── postgres
│   │   │   └── postgres.go
│   │   └── repository.go
│   ├── routers
│   │   └── links.go
│   ├── service
│   │   ├── errors.go
│   │   └── links.go
│   └── structs
│       └── structs.go
└── pkg
    └── utils
        ├── errors.go
        └── utils.go
```

#### `cmd` package

Entrypoint of application with `main.go`

#### `internal` package

(See the [Go 1.4 release notes](https://golang.org/doc/go1.4#internalpackages)) 

Application internal main logic

- `config` configuration loading modules
- `handlers` request handlers for routers
- `models` ORM models
- `repository` storage abstraction (see the [Repository design pattern](https://medium.com/@pererikbergman/repository-design-pattern-e28c0f3e4a30))
  - `memory` in-memory implementation
  - `postgres` PostgreSQL implementation
- `routers` Gin API routers
- `service` business logic
- `structs` request/response models

#### `pgk` package

Exportable packages like `utils` with link code generating function

#### `api` package

API layout related logic like Swagger setup

#### `docs` package

Documentation related specs (Swagger)

## References

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/)
- [Best Practices of Building Web Apps with Gin & Golang](https://www.squash.io/optimizing-gin-in-golang-project-structuring-error-handling-and-testing/)
- Templates
    - [Go RESTful API template](https://github.com/gilcrest/diygoapi)

## Go ahead

- [Откажитесь уже наконец от gin, echo и <иной ваш фреймворк>](https://habr.com/ru/companies/ozonbank/articles/817381/)