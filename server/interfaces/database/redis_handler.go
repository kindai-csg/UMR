package database

type RedisHandler interface {
	Set(string, string) (error) 
	Get(string) (string, error)
	RPush(string, []string) (error) 
	LPop(string, int)  ([]string, error)
}
