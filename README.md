# Golang Bootcamp

## Introduction

This project consists on a REST API which works based in the [Todoist API](https://developer.todoist.com/rest/v1/#overview).

## Requirements

1. You need at least `Go 1.14` installed to run this project.

## Installation

1. After cloning the project you have to build the project with the following command 
```shell
go build
```

1. Once the project is build you have to create a `.env` file, you can use the `.env.example` file as a template for this file. You can follow the guide of the [Todoist API](https://developer.todoist.com/rest/v1/#overview) to setup your API token.

1. For running the project you can use the following command in the root directory of the project:
```shell
go run main.go
```

1. To run the tests you can use command
```shell
go test ./...
```
