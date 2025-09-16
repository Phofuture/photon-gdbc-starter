package gdbc

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/Phofuture/photon-core-starter/log"

	"gorm.io/gorm"
)

type DbAction func(ctx context.Context, db *gorm.DB) (err error)
type entry struct {
	dbPointer **gorm.DB
	name      string
}

var customAction []DbAction
var dbMap sync.Map
var primaryDb *gorm.DB
var registeredDbs []entry

func RegisterDbCustomize(action DbAction) {
	customAction = append(customAction, action)
}

func RegisterDb(db **gorm.DB, names ...string) {
	if db == nil {
		panic("db pointer is nil")
	}
	name := "primary"
	if len(names) > 0 {
		name = names[0]
	}
	registeredDbs = append(registeredDbs, entry{
		dbPointer: db,
		name:      name,
	})
}

func GetGDBCTemplate() *gorm.DB {
	return primaryDb
}

func GetGDBCTemplateByName(name string) *gorm.DB {
	if db, ok := dbMap.Load(name); ok {
		if gdb, valid := db.(*gorm.DB); valid {
			return gdb
		}
	}
	return nil
}

func Start(ctx context.Context) (err error) {
	log.Logger().Info(ctx, "init database")
	if err != nil {
		log.Logger().Error(ctx, "failed to get database config", "error", err)
		return
	}

	if len(databaseConfig.Database.DataSources) == 0 {
		if primaryDb, err = setDatasource(ctx, "primary", databaseConfig.Database.Primary); err != nil {
			return
		}
	}

	for name, connectData := range databaseConfig.Database.DataSources {
		var db *gorm.DB
		if db, err = setDatasource(ctx, name, connectData); err != nil {
			return
		}
		dbMap.Store(name, db)
		if connectData.IfPrimary {
			primaryDb = db
		}
	}

	if primaryDb == nil {
		return fmt.Errorf("no primary database configured")
	}

	for _, entry := range registeredDbs {

		if entry.name == "primary" {
			*(entry.dbPointer) = primaryDb
			continue
		}

		db, ok := dbMap.Load(entry.name)
		if !ok {
			return fmt.Errorf("no db found for name %s", entry.name)
		}

		gdb, valid := db.(*gorm.DB)
		if !valid {
			return fmt.Errorf("invalid db type for name %s", entry.name)
		}
		*(entry.dbPointer) = gdb
	}

	return
}

func setDatasource(ctx context.Context, name string, connectData ConnectData) (db *gorm.DB, err error) {
	if db, err = connect(ctx, connectData); err != nil {
		log.Logger().Error(ctx, "fail to connect master database", "error", err, "config", databaseConfig)
		return
	}
	for _, action := range customAction {
		if err = action(ctx, db); err != nil {
			msg := fmt.Sprintf("failed to customize database : %s", name)
			log.Logger().Error(ctx, msg, "error", err)
			return
		}
	}
	return
}

// 連線資料庫
func connect(ctx context.Context, connectData ConnectData) (db *gorm.DB, err error) {

	dialector := NewDialector(connectData)
	if db, err = gorm.Open(dialector); err != nil {
		log.Logger().Error(ctx, "failed to connect database", "error", err, "datasource", connectData)
		return
	}

	err = setConnectPool(ctx, db, connectData)
	return
}

// 設定 connection pool
func setConnectPool(ctx context.Context, db *gorm.DB, connectData ConnectData) (err error) {
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		log.Logger().Error(ctx, "failed to init database", "error", err)
		return
	}

	maxIdleConns := connectData.Connection.MaxIdleConns
	if maxIdleConns == 0 {
		maxIdleConns = 10
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)

	maxOpenConns := connectData.Connection.MaxOpenConns
	if maxOpenConns == 0 {
		maxOpenConns = 50
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)

	maxLifetime := connectData.Connection.MaxLifetimeSecond
	if maxLifetime == 0 {
		maxLifetime = 600
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	return
}
