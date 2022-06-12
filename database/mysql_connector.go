package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	config *MySqlConfig
	db     *gorm.DB
}

type MySqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}

var _ Connector = (*MySQLConnector)(nil)

func NewMySQLConnector(config *MySqlConfig) Connector {
	c := &MySQLConnector{}

	c.initConfig(config)
	c.Connect()

	return c
}

// Connect creates a new gorm.DB instance.
func (c *MySQLConnector) Connect() {
	db, err := gorm.Open(mysql.Open(c.parseDSN()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	c.db = db
}

// initConfig initializes the config.
func (c *MySQLConnector) initConfig(config *MySqlConfig) {
	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == 0 {
		config.Port = 3306
	}

	if config.Charset == "" {
		config.Charset = "utf8mb4"
	}

	c.config = config
}

func (c *MySQLConnector) parseDSN() string {
	// format "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.config.Username, c.config.Password, c.config.Host, c.config.Port, c.config.Database, c.config.Charset)
}

func (c *MySQLConnector) DB() interface{} {
	return c.db
}
