package infrastructure_test

import (
	inf "github.com/kindaidensan/UMR/infrastructure"
	"testing"
)

func TestNewLdapHandler(t *testing.T) {
	ldapHandler := inf.NewLdapHandler()
	if ldapHandler == nil {
		t.Errorf("faild: cant create LdapHandler instance")
	}
}

func _TestLdapAddRequest(t *testing.T) {
	ldapHandler := inf.NewLdapHandler()
	if ldapHandler == nil {
		t.Errorf("faild: cant create LdapHandler instance")
		return
	}
	request := []string {
		"dn", "cn=test,ou=member,ou=account,dc=kindai-csg,dc=dev",
		"objectClass", "posixAccount",
		"objectClass", "inetOrgPerson",
		"cn", "test",
		"uid", "1400300001", 
		"uidNumber", "2000",
		"gidNumber", "1002",
		"homeDirectory", "/home/test",
		"userPassword", "testPassword",
		"sn", "Test",
		"displayName", "Test",
		"mail", "test@hogehoge.com",
	}
	err := ldapHandler.AddRequest(request)
	if err != nil {
		t.Errorf("faild: err request")
		t.Errorf(err.Error())
	}
}

func TestLdapSearchRequest(t *testing.T) {
	ldapHandler := inf.NewLdapHandler()
	if ldapHandler == nil {
		t.Errorf("faild: cant create LdapHandler instance")
		return
	}
	err := ldapHandler.SearchRequest("test")
	if err != nil {
		t.Errorf("faild: err request")
		t.Errorf(err.Error())
	}
	

}