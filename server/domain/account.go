package domain

type Account struct {
	ID string
	Password string
	Name string
	EmailAddress string
	AccountType string
}

type Accounts []Account
