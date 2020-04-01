package infrastructure

import (
	"gopkg.in/ldap.v2"
)

type LdapHandler struct {
	connection *ldap.Conn 
}

func NewLdapHandler() *LdapHandler{
	ldap, err := ldap.Dial("tcp", "ldap:389")
	if err != nil {
		return nil 
	}
	ldapHandler := LdapHandler {
		connection: ldap, 
	}
	err = ldapHandler.Bind()
	if err != nil {
		return nil
	}
	return &ldapHandler
}

func (handler *LdapHandler) Bind() error {
	err := handler.connection.Bind("cn=Manager,dc=kindai-csg,dc=dev", "densan")
	if  err != nil {
		return err
	}
	return nil
}

func (handler *LdapHandler) AddRequest(request []string) error {
	addRequest := ldap.NewAddRequest(request[1])
	objectClass := []string{}
	for i, str := range request {
		if str == "objectClass" {
			objectClass = append(objectClass, request[i+1])
		}
	}
	addRequest.Attribute("objectClass", objectClass)
	for i := 2; i < len(request); i += 2 {
		if request[i] == "objectClass" {
			continue
		}
		addRequest.Attribute(request[i], []string{ request[i+1] })
	}
	err := handler.connection.Add(addRequest)
	if err != nil {
		return err
	}
	return nil
}

func (handler *LdapHandler) UpdateRequest(id string, request []string) error {
	return nil
}

func (handler *LdapHandler) DeleteRequest(cn string) error {
	delRequest := ldap.NewDelRequest(cn, nil)
	err := handler.connection.Del(delRequest)
	if err != nil {
		return err
	}
	return nil
}

func (handler *LdapHandler) SearchRequest(id string, attributes []string) ([][]string, error) {
	searchRequest := ldap.NewSearchRequest(
		"ou=account,dc=kindai-csg,dc=dev",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(cn="+id+")",
		attributes,
		nil,
	)
	result, err := handler.connection.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	resultArray := [][]string{}
	for i, entry := range result.Entries {
		resultArray = append(resultArray, []string{})
		for _, attr := range attributes {
			if attr == "dn" {
				resultArray[i] = append(resultArray[i], entry.DN)
				continue
			} 
			resultArray[i] = append(resultArray[i], entry.GetAttributeValue(attr))	
		}
	}
	return resultArray, nil
}

func (handler *LdapHandler) Close() {
	handler.connection.Close()
}