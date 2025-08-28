package gdbc

import "gorm.io/gorm"

func NewDialector(data ConnectData) gorm.Dialector {
	switch data.Driver {
	case "mysql":
		return newMySQLDialector(data)
	case "postgres":
		return newPostgresDialector(data)
	}
	return nil
}
