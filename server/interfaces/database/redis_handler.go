package database

type RedisHandler interface {
	Set(string, string) (error) 
	Get(string) (string, error)
	ExpireSetKey(string, int) (error)
	GetTtl(string) (int, error)
	RPush(string, []string) (error) 
	LPop(string, int)  ([]string, error)
}
