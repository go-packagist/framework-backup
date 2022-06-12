package facades

import (
	"errors"
	"github.com/go-packagist/framework/database"
)

// Database returns the database manager.
func Database() (*database.Manager, error) {
	db, err := App().Make("database")
	if err != nil {
		return nil, err
	}

	switch db.(type) {
	case *database.Manager:
		return db.(*database.Manager), nil
	default:
		return nil, errors.New("database is not a database manager")
	}
}

// MustDatabase returns the database manager.
func MustDatabase() *database.Manager {
	db, _ := Database()

	return db
}
