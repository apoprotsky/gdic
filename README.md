# Go Service Container

[![Go Report Card](https://goreportcard.com/badge/github.com/apoprotsky/gservices)](https://goreportcard.com/report/github.com/apoprotsky/gservices)

`gservices` library implements `Service Container` pattern.
It creates instances of objects, resolving their dependencies using simple API.

To resolve services dependencies simply add a public Pointer type field in your service structure or
declare method `Init` which arguments is pointers to other services.

## Usage example

Declaring `db` package

```go
package db

type Service struct {}
```

Declaring `accounts` package

```go
package accounts

import "app/db"

type Account struct {/*...*/}

type Service struct {
    DBService *db.Service
}

func (svc *Service) Get(/*...*/) *Account {
    var data = svc.DBService.Find(/*...*/)
    /*...*/
    return &Account{data: data}
}
```

Main application package

```go
package main

import (
    "app/accounts"
    "github.com/apoprotsky/gservices"
)

func main() {
    accountsService := gservices.Get((*accounts.Service)(nil)).(*accounts.Service)
    account := accountsService.Get(/*...*/)
    /*...*/
}
```

See example directory for more information.
