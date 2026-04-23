package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type Account struct {
	login    string
	password string
	url      string
}

func NewAccount(login, password, userURL string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID LOGIN")
	}
	_, err := url.ParseRequestURI(userURL)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}
	res := &Account{
		login:    login,
		password: password,
		url:      userURL,
	}
	if res.password == "" {
		res.generatePassword(5)
	}
	return res, nil

}

func (acc *Account) generatePassword(n int) {
	letter := []rune("1234567890-=asdfghjkl;'" +
		"zxcvbnm,./")
	res := make([]rune, n)
	for item := range res {
		res[item] = letter[rand.IntN(len(letter))]
	}
	acc.password = string(res)
}

func (acc *Account) OutputDataAccount() {
	fmt.Printf("Login: %s | Password: %s | URL: %s\n",
		acc.login, acc.password, acc.url)
}
