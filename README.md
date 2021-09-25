## (WIP) Sandbox REST API with Golang

- Session authentication with postgres using cookie (not JWT)


## Notes

### Requirements

```bash
brew install golang-migrate
go install github.com/cespare/reflex@latest
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```

### How to Run

```bash
make postgres
make resetdb
make devw
```

### Create Migration File

```bash
migrate create -ext sql --dir db/migrations -seq NAME
```

## References

- https://github.com/techschool/simplebank
- https://github.com/ardanlabs/service
- https://github.com/DeNA/codelabs/tree/master/sources/testable-architecture-with-go
