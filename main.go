package main

import (
	"fmt"
	"password/account"
)

func main() {
	fmt.Println("o")
	login := promptData("Login")
	password := promptData("Password")
	userURL := promptData("URL")

	acc, err := account.NewAccount(login, password, userURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	acc.OutputDataAccount()
}

func promptData(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}
