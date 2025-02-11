## Implementation solutions

### Mutual uniqueness of shortened and original url

### Memory implementation

A bidirectional hashmap was used for the solution, as it allows checking for the existence of a link both from the
original url and from the shortened one code in the optimal time.

```go
type LinkRepository struct {
	linksByCode map[string]*models.Link
	codesByUrl  map[string]string
	...
}
```

### PostgreSQL implementation

In case of database implementation, two unique indexes are used for the links table (original url, shortened link code),
which allows to achieve the same effect.

```go
type Link struct {
	...
	ShortCode string `gorm:"uniqueIndex;not null"`
	URL       string `gorm:"uniqueIndex;not null"`
}
```

## References

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/)
- [Best Practices of Building Web Apps with Gin & Golang](https://www.squash.io/optimizing-gin-in-golang-project-structuring-error-handling-and-testing/)
- Templates
    - [Go RESTful API template](https://github.com/gilcrest/diygoapi)

## Go ahead

- [Откажитесь уже наконец от gin, echo и <иной ваш фреймворк>](https://habr.com/ru/companies/ozonbank/articles/817381/)