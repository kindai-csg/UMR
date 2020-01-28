package domain

type RegularAccount struct {
	ID string `validate:"required,min=4,max=15,alphanumunicode"`
	FirstName string `validate:"required"`
	LastName string `validate:"required"`
	FirstNameKana string `validate:"required"`
	LastNameKana string `validate:"required"`
	StudentNumber string `validate:"required"`
	Department string `validate:"required"`
	Grade int `validate:"required,min=1,max=4"`
	PostalCode string `validate:"required"`
	StreetAddress string `validate:"required"`
	PhoneNumber string `validate:"required"`
	EmergencyPhoneNumber string `validate:"required"`
}

type RegularAccounts []RegularAccount
