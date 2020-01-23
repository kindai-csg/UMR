package domain

type UserData struct {
	ID string
	UserType string
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

type UserDatas []UserData
