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

## Dependencies

I'm not writing all of these from scratch, big thanks to these great libraries and their contributors.

### CLI/Executable Dependencies

Dependencies to other projects, can be installed using [Brew](https://brew.sh/):

- Install [Migrate](https://github.com/golang-migrate/migrate) for Go:
```
brew install golang-migrate
```

### Libraries

- [echo](https://github.com/labstack/echo)
- [pgx](https://github.com/jackc/pgx)
- [godotenv](https://github.com/joho/godotenv)
- [faker](https://github.com/go-faker/faker)