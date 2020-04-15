package database

import (
	"github.com/kindaidensan/UMR/domain" 
	"strings"
	// "database/sql"
	"crypto/md5"
	// "encoding/hex"
	"encoding/base64"
)

type AccountRepository struct {
	LdapHandler LdapHandler
	RedisHandler RedisHandler
	SqlHandler SqlHandler
}

func NewAccountRepository(ldapHandler LdapHandler, redisHandler RedisHandler, sqlHandler SqlHandler) *AccountRepository {
	accountRepository := AccountRepository {
		LdapHandler: ldapHandler,
		RedisHandler: redisHandler,
		SqlHandler: sqlHandler,
	}
	return &accountRepository
}

func (repo *AccountRepository) TemporaryStore(account domain.Account) error {
	err := repo.RedisHandler.RPush("tmp_"+account.ID, []string{account.Password, account.Name, account.EmailAddress, account.StudentNumber, account.AccountType})
	if err != nil {
		return err
	}
	err = repo.RedisHandler.ExpireKey("tmp_"+account.ID, 1800)
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
	md5 := md5.Sum([]byte(account.Password))
	password := base64.StdEncoding.EncodeToString(md5[:])
	request := []string {
		"dn", "cn="+account.ID+","+ou+",dc=kindai-csg,dc=dev",
		"objectClass", "posixAccount",
		"objectClass", "inetOrgPerson",
		"cn", account.ID,
		"uid", account.StudentNumber,
		"uidNumber", account.UserIdNumber,
		"gidNumber", account.GroupIdNumber,
		"homeDirectory", account.HomeDirectory,
		"userPassword", "{MD5}"+password,
		"sn", account.Name,
		"displayName", account.Name,
		"mail", account.EmailAddress,
	}
	err := repo.LdapHandler.AddRequest(request)
	if   err != nil {
		return err
	}
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

func (repo *AccountRepository) GetAllAccounts() ([]domain.Account, error) {
	account_data, err := repo.LdapHandler.SearchRequest("*", []string { "cn", "displayName", "mail", "uid", "dn" })
	if err != nil {
		return nil, err
	}
	accounts := []domain.Account {}
	for _, account := range account_data {
		accounts = append(accounts, domain.Account {
			ID: account[0],
			Name: account[1],
			EmailAddress: account[2],
			StudentNumber: account[3],
			AccountType: account[4],
		})
	}
	return accounts, nil
}

func (repo *AccountRepository) GetAllNonActiveAccountID() ([]string, error) {
	tmps, err := repo.RedisHandler.GetKeys("tmp_*")
	if err != nil {
		return nil, err
	}
	auths, err := repo.RedisHandler.GetKeys("auth_*")
	if err != nil {
		return nil, err
	}
	accounts := []string {}
	for _, tmp := range tmps {
		id := strings.Replace(tmp, "tmp_", "", 1)
		isAuth := false
		for _, auth := range auths {
			if id == strings.Replace(auth, "auth_", "", 1) {
				isAuth = true
				break
			} 
		}
		if !isAuth {
			accounts = append(accounts, id)
		}
	}
	return accounts, nil
} 

func (repo *AccountRepository) DeleteAccount(id string) error {
	err := repo.LdapHandler.DeleteRequest("cn="+id+",ou=account,dc=kindai-csg,dc=dev")
	if err != nil {
		return err
	}
	return nil
}

func (repo *AccountRepository) GetAdminAccounts() ([]domain.AdminAccount, error) {
	rows, err := repo.SqlHandler.Query("select * from admin")
	if err != nil {
		return nil, err
	}

	accounts := []domain.AdminAccount {}
	for rows.Next() {
		var id string 
		var password string
		if err := rows.Scan(&id, &password); err != nil {
			return nil, err
		}
		accounts = append(accounts, domain.AdminAccount {
			ID: id,
			Password: password,
		})
	}
	return accounts, nil 
}
