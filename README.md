# GINGON BOILERPLATE

## Requirement tools
-   [Golang migrate](https://github.com/golang-migrate/migrate) => migration management. Installation guide: [windows](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-windows/), [linux ubuntu](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/), [mac os](https://formulae.brew.sh/formula/golang-migrate)

## Commands
### Migration
⌨️   migrate command example for postgresql database

```migrate -database "postgres://postgres:@localhost:5432/gin_gonic?sslmode=disable" -path database/migrations up```

```migrate -database "postgres://postgres:@localhost:5432/gin_gonic?sslmode=disable" -path database/migrations down```
-   for other databases, follow documentation [here](https://github.com/golang-migrate/migrate#databases)