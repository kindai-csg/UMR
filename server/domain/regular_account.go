package domain

type RegularAccount struct {
	ID string
	FirstName string
	LastName string
	FirstNameKana string
	LastNameKana string
	StudentNumber string
	Department string
	Grade int
	PostalCode string
	StreetAddress string
	PhoneNumber string
	EmergencyPhoneNumber string
}

type RegularAccounts []RegularAccount
