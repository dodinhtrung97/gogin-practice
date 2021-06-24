# GOGIN PRACTICE

## Service
Service is ran on port `8080`

Database configurations are sorted in `\config\db_config.json` and read in `animal-repository` using `gonfig` (Local environment)

Database connection and ORM are handled by `GORM`

## What's in it
* Practice creating HTTP requests with GoGin
* Connect to MySQL Database using `GORM`
* Handle database connection and ORM using `GORM`
* Practice GoLang HTTP error handling
* Practice GoLang struct tag and GORM's entity tag
* One-to-one relationship
* Cascading
* Packaging

## Environment
Install [GoLang](https://golang.org/doc/install)

## Run It
Each directory should be ran as seperated projects

To initialize module

```
$ go mod init {module-name}
```

To install module's dependencies

```
$ go mod tidy
```