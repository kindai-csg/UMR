package database

type LdapHandler interface {
	AddRequest([]string) (error) 
	UpdateRequest(string, []string) (error) 
	DeleteRequest(string) (error) 
	SearchRequest(string) (error) 
}
