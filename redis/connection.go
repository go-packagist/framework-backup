package redis

import "time"

type Connection interface {
	Connect(map[string]interface{})
	Set(string, interface{}, time.Duration) (string, error)
	Get(string) (string, error)
	Echo(string) (string, error)
	Ping() (string, error)
	Del(...string) (int64, error)
	Exists(...string) (int64, error)
}
