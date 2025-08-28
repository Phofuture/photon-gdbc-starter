package gdbc

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresDialector(connectData ConnectData) gorm.Dialector {
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connectData.Host,
		connectData.Port,
		connectData.Username,
		connectData.Password,
		connectData.Schema,
	)
	
	return postgres.Open(dsn)
}
