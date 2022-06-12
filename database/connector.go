package database

type Connector interface {
	Connect()
	DB() interface{}
}
