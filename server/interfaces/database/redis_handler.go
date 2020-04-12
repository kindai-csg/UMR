package database

type RedisHandler interface {
	Set(string, string) (error) 
	Get(string) (string, error)
	Del(string) (error)
	MultiDel([]string) (error)
	ExpireSetKey(string, string, int) (error)
	ExpireKey(string, int) (error)
	GetTtl(string) (int, error)
	RPush(string, []string) (error) 
	LPop(string, int)  ([]string, error)
	GetKeys(string) ([]string, error)
	Incr(string) (int, error)
}
