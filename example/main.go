package main

import (
	"github.com/apoprotsky/gservices/example/accounts"
)

func main() {
	accountsService := accounts.GetService()
	account := accountsService.GetByName("admin")
	println("admin email: " + account.Email)
}
