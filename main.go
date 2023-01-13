package main

import (
	"github.com/91savage/learngo/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("sehun")
	fmt.Println(account)
}