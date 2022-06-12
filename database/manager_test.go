package database

import (
	"testing"
)

func createConfig() *Config {
	return &Config{
		Default: "mysql",
		Connections: map[string]interface{}{
			"mysql": &MySqlConfig{
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "123456",
				Database: "test",
			},
		},
	}
}

func createManager() *Manager {
	return NewManager(createConfig())
}

func TestManager_Connection(t *testing.T) {
	// m := createManager()
	//
	// type User struct {
	// 	ID       int    `gorm:"primary_key"`
	// 	Username string `gorm:"column:username"`
	// }
	//
	// var user User
	// m.Connection().DB().(*gorm.DB).First(&user)
	//
	// fmt.Println(user)
}
