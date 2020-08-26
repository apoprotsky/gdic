# Go Dependency Injection Container

[![Go Report Card](https://goreportcard.com/badge/github.com/apoprotsky/gdic)](https://goreportcard.com/report/github.com/apoprotsky/gdic)

`gdic` library implements `Dependency Injection Container` pattern also known as `Service Container`.
It creates instances of objects, resolving their dependencies using simple API.

## Usage example

Declaring `db` package

```go
package db

type Service struct {}

func (svc *Service) New(data string) int {/*...*/}

func New() *Service {
    return &Service{}
}
```

Declaring `accounts` package

```go
package accounts

import "app/db"

type Account struct {/*...*/}

type Service struct {
    dbService *db.Service
}

func (svc *Service) Get(/*...*/) *Account {
    var data = svc.dbService.Find(/*...*/)
    /*...*/
    return &Account{data: data}
}

func New(dbService *db.Service) *Service {
    return &Service{dbService: dbService}
}
```

Main application package

```go
package main

import (
    "app/accounts"
    "app/db"
    "github.com/apoprotsky/gdic"
)

func main() {
    // Init GDIC
    gdic.RegisterProviderOrPanic(db.New)
    gdic.RegisterProviderOrPanic(accounts.New)
    // Use GDIC
    var accountsService = gdic.GetOrPanic((*users.Service)(nil)).(*users.Service)
    var account = accountsService.Get(/*...*/)
    /*...*/
}
```
