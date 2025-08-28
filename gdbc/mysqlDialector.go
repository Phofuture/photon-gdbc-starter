package gdbc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMySQLDialector(connectData ConnectData) gorm.Dialector {
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		connectData.Username,
		connectData.Password,
		connectData.Host,
		connectData.Port,
		connectData.Schema)
	
	return mysql.Open(dsn)
}
