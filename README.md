# Kwik-E-Mart Demo App <!-- omit in toc -->

![image](https://user-images.githubusercontent.com/1924063/140587139-4e42d91a-1db4-4bc6-a3a8-c66874d56f08.png)

A simple Golang-based application that queries a PostgreSQL database named `kwikemart` to read and return customer data stored in the `customers` table.

- [Database Structure](#database-structure)
- [Required Environment Variables](#required-environment-variables)
- [Run with Summon](#run-with-summon)
  - [Pre-Requisites](#pre-requisites)
  - [Command](#command)
- [Build](#build)
  - [Binary From Source](#binary-from-source)
    - [Build for current OS and OS Architecture](#build-for-current-os-and-os-architecture)
    - [Build for different OS and OS Architecture](#build-for-different-os-and-os-architecture)
  - [Build From Dockerfile](#build-from-dockerfile)
    - [Run in Docker](#run-in-docker)
- [Example Output](#example-output)
- [License](#license)

## Database Structure

```shell
$ make create_table
$ make insert_customers
```

Database Name: `kwikemart`

Table Name: `customers`

|id|first_name|last_name|pmt_type|
|-|-|-|-|
|1|Homer|Simpson|cash|
|2|Montgomery|Burns|credit|
|3|Barney|Gumble|debit|
|4|Waylon|Smithers|cash|
|5|Ned|Flanders|credit|

## Required Environment Variables

Provided for [Summon](https://cyberark.github.io/summon) in [secrets.yml](secrets.yml).

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

OR

```shell
make run
```

## Build

### Binary From Source

#### Build for current OS and OS Architecture

```shell
go build -o bin/kwikemart .
```

#### Build for different OS and OS Architecture

```shell
make compile
```

### Build From Dockerfile

```shell
make build
```

#### Run in Docker

```shell
summon -p summon-conjur docker run --name kwik-e-mart -d \
  --restart unless-stopped \
  --env-file @SUMMONENVFILE \
  -p 8080:8080
  nfmsjoeg/kwik-e-mart:latest
```

OR

```shell
make docker_run
```

## Example Output

```shell
$ open http://localhost:8080
```

```html
-------------------------------------------------------
Connected successfully to conjur-demo.xxxxxx.rds.amazonaws.com
Database Username: xxxxxx
Database Password: xxxxxx
-------------------------------------------------------

id             first_name     last_name      pmt_type
-------------------------------------------------------
1              Homer          Simpson        cash
2              Montgomery     Burns          credit
3              Barney         Gumble         debit
4              Waylon         Smithers       cash
5              Ned            Flanders       credit

```

## License

[MIT](LICENSE)
