package domain_test

import(
	"github.com/kindaidensan/UMR/domain"
	"testing"
	"gopkg.in/go-playground/validator.v9"
)

func TestAccountValidation(t *testing.T) {
	validate := validator.New()

	account := &domain.Account{
		ID: "test",
		Password: "12345678",
		Name: "aiueo",
		EmailAddress: "test@hogehoge.com",
		StudentNumber: "1000000000",
		AccountType: "Regular",
	}

	// ---------------------------ID---------------------------
	account.ID = "t"
	err := validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.ID = "test123456789012"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.ID = "test@test"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.ID = "test_test1"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.ID = "テスト"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.ID = "four"
	err = validate.Struct(account)
	if err != nil {
		t.Errorf("Faild validation")
	}

	account.ID = "15mojiIDdasuyo1"
	err = validate.Struct(account)
	if err != nil {
		t.Errorf("Faild validation")
	}
	// ---------------------------Password---------------------------
	account.ID = "test"
	account.Password = "test"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.Password = "テストだ"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.Password = "テストだテスト"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.Password = "12345678"
	err = validate.Struct(account)
	if err != nil {
		t.Errorf("Faild validation")
	}
	// ---------------------------name---------------------------
	account.Password = "123456789"
	account.Name = ""
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}
	// ---------------------------email---------------------------
	account.Name = "aiueo"
	account.EmailAddress = "test@test"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}

	account.EmailAddress = "testt.com"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}
	// ---------------------------account type---------------------------
	account.EmailAddress = "test@hogehoge.com"
	account.AccountType = "re"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}
	// ---------------------------studnet number---------------------------
	account.AccountType = "Regular"
	account.StudentNumber = "100000000a"
	err = validate.Struct(account)
	if err == nil {
		t.Errorf("Faild validation")
	}
}