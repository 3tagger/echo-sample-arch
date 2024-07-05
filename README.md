# echo-sample-arch

Sample of simple web server built using Echo framework.

## Goal of This Project

I want to explore the [Echo](https://echo.labstack.com/) framework, one of the popular Go framework.

In this current state, it only shows the bare minimum usage of Echo as a web server.

## Details

### Run the Project

To run this project, simply run this from the root of the project:
```
make run
```

### Run the Dependencies

To run the local developemnt environment dependencies, such as PostgreSQL database:
```
make up
```

## Dependencies

Dependencies to other projects, can be installed using [Brew](https://brew.sh/):

- Install [Migrate](https://github.com/golang-migrate/migrate) for Go:
```
brew install golang-migrate
```