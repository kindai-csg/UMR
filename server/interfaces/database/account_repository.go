package database

import (
	"github.com/kindaidensan/UMR/domain" 
)

type AccountRepository struct {
	LdapHandler LdapHandler
	RedisHandler RedisHandler
}

func NewAccountRepository(ldapHandler LdapHandler, redisHandler RedisHandler) *AccountRepository {
	accountRepository := AccountRepository {
		LdapHandler: ldapHandler,
		RedisHandler: redisHandler,
	}
	return &accountRepository
}

func (repo *AccountRepository) TemporaryStore(account domain.Account) error {
	err := repo.RedisHandler.RPush(account.ID, []string{account.Password, account.Name, account.EmailAddress, account.StudentNumber, account.AccountType})
	if err != nil {
		return err
	}
	return nil
}

func (repo *AccountRepository) FindByIdFromTemporary(id string) (domain.Account, error) {
	account, err := repo.RedisHandler.LPop(id, 5)
	if err != nil {
		return domain.Account{}, err
	}
	return domain.Account{ID: id, Password: account[0], Name: account[1], EmailAddress: account[2], StudentNumber: account[3], AccountType: account[4]}, nil
}

func (repo *AccountRepository) Store(account domain.Account) error {
	ou := ""
	if account.AccountType == "Regular" {
		ou = "ou=member,ou=account"
	} else {
		ou = "ou=account"
	}
	request := []string {
		"dn", "cn="+account.ID+","+ou+",dc=kindai-csg,dc=dev",
		"objectClass", "posixAccount",
		"objectClass", "inetOrgPerson",
		"cn", account.ID,
		"uid", account.StudentNumber,
		"uidNumber", account.StudentNumber[:2]+account.StudentNumber[7:],
		"gidNumber", "1002",
		"homeDirectory", "/home/"+account.ID,
		"userPassword", account.Password,
		"displayName", account.Name,
		"mail", account.EmailAddress,
	}
	err := repo.LdapHandler.AddRequest(request)
	if   err != nil {
		return err
	}
	return nil
}
