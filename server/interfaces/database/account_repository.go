package database

import (
	"github.com/kindaidensan/UMR/domain" 
	"strings"
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
	err := repo.RedisHandler.RPush("tmp_"+account.ID, []string{account.Password, account.Name, account.EmailAddress, account.StudentNumber, account.AccountType})
	if err != nil {
		return err
	}
	return nil
}

func (repo *AccountRepository) FindByIdFromTemporary(id string) (domain.Account, error) {
	account, err := repo.RedisHandler.LPop("tmp_"+id, 5)
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
		"uidNumber", account.UserIdNumber,
		"gidNumber", account.GroupIdNumber,
		"homeDirectory", account.HomeDirectory,
		"userPassword", account.Password,
		"sn", account.Name,
		"displayName", account.Name,
		"mail", account.EmailAddress,
	}
	err := repo.LdapHandler.AddRequest(request)
	if   err != nil {
		return err
	}
	_ = repo.RedisHandler.MultiDel([]string { "auth_"+account.ID, "count_"+account.ID  })
	return nil
}

func (repo *AccountRepository) GetAllUserID() ([]string, error) {
	tmpIds, err := repo.RedisHandler.GetKeys("tmp_*")
	if err != nil {
		return nil, err
	}
	for i, _ := range tmpIds {
		tmpIds[i] = strings.Replace(tmpIds[i], "tmp_", "", 1) 
	}
	ids, err := repo.LdapHandler.SearchRequest("*", []string { "cn" })
	for _, cn := range ids {
		tmpIds = append(tmpIds, cn[0])
	}
	return tmpIds, nil
}
