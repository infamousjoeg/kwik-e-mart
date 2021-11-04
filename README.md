# Qwik-E-Mart Demo App <!-- omit in toc -->

A simple Golang-based application that queries a PostgreSQL database named `qwikemart` to read and return customer data stored in the `customers` table.

- [Database Structure](#database-structure)
- [Required Environment Variables](#required-environment-variables)
- [Run with Summon](#run-with-summon)
  - [Pre-Requisites](#pre-requisites)
  - [Command](#command)
- [Build Binary From Source](#build-binary-from-source)
  - [Build for current OS and OS Architecture](#build-for-current-os-and-os-architecture)
  - [Build for different OS and OS Architecture](#build-for-different-os-and-os-architecture)
- [License](#license)

## Database Structure

Database Name: `qwikemart`
Table Name: `customers`

|id|first_name|last_name|pmt_type|
|-|-|-|-|
|1|Homer|Simpson|cash|
|2|Montgomery|Burns|credit|
|3|Barney|Gumble|debit|
|4|Waylon|Smithers|cash|
|5|Ned|Flanders|credit|

## Required Environment Variables

|variable name|expected value|
|-|-|
|DB_HOST|hostname to PostgreSQL database|
|DB_USERNAME|username to use for authentication|
|DB_PASSWORD|password to use for authentication|
|DB_NAME|name of the database to connect to|

## Run with Summon

### Pre-Requisites

* [Summon](https://cyberark.github.io/summon)
* [Supported Summon Provider](https://cyberark.github.io/summon/#providers)
* [secrets.yml](secrets.yml) file

### Command

```shell
summon -p $provider_name go run main.go
```

```shell
summon -p $provider_name -e $env_name go run main.go
```

## Build Binary From Source

### Build for current OS and OS Architecture

```shell
go build .
```

### Build for different OS and OS Architecture

```shell
GOOS=windows GOARCH=amd64 go build .
```

* `GOOS` is the operating system name
* `GOARCH` is the architecture to compile for

## License

[MIT](LICENSE)