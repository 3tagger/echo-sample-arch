# echo-sample-arch

Sample of simple web server built using Echo framework.

## Goal of This Project

I want to explore the [Echo](https://echo.labstack.com/) framework, one of the popular Go framework.

In this current state, it only shows the bare minimum usage of Echo as a web server.

## Details

If you run it first time, run these `make` targets in order.

### Run the Dependencies

To run the local developemnt environment dependencies, such as PostgreSQL database:
```
make up
```

### Run the Migration

To run the migration, install [Migrate](https://github.com/golang-migrate/migrate)
```
make migrate-up
```

### Run the Seeder

This is optional, run only when you need sample data to use this API. Warning: don't run this on production DB.
```
make seed
```

### Run the Project

To run this project, simply run this from the root of the project:
```
make run
```

### Generate Documentation (Swagger)

To generate the documentation for the API, run this:
```
make gen-doc
```

## Dependencies

I'm not writing all of these from scratch, big thanks to these great libraries and their contributors.

This project uses Go `v1.22.x`.

### CLI/Executable Dependencies

Dependencies to other projects, can be installed using [Brew](https://brew.sh/):

- Install [Migrate](https://github.com/golang-migrate/migrate) for Go:
```
brew install golang-migrate
```

- Install [httpyac](https://httpyac.github.io/) VSCode extension for integrated HTTP call with VSCode. By adding this extension, you can run the `./request/*.http` files directly from your browser.

### Libraries

- [echo](https://github.com/labstack/echo)
- [pgx](https://github.com/jackc/pgx)
- [godotenv](https://github.com/joho/godotenv)
- [faker](https://github.com/go-faker/faker)
- [mockery](https://github.com/vektra/mockery)
- [swaggo](https://github.com/swaggo/swag)
- [validator](https://github.com/go-playground/validator)
