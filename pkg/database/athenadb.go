package database

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

type AthenaDb interface {
	Connect()
	Disconnect()
}

type AthenaDbConfig struct {
	ConnectionString string
}

type AthenaMySqlDb struct {
	Config AthenaDbConfig
	Db     *gorm.DB
}

func NewAthenaMySqlDb(c AthenaDbConfig) *AthenaMySqlDb {
	return &AthenaMySqlDb{
		Config: c,
	}
}

func (a *AthenaMySqlDb) Connect() {
	db, err := gorm.Open(mysql.Open(a.Config.ConnectionString), &gorm.Config{})

	if err != nil {
		slog.Error("Couldn't connect to MySQL DB", "error", err)
		panic("Could not connect to MySQL DB")
	}

	a.Db = db
}

func (a *AthenaMySqlDb) Disconnect() {
	db, err := a.Db.DB()
	if err != nil {
		slog.Error("Couldn't find DB to disconnect", "error", err)
	}

	err = db.Close()
	if err != nil {
		slog.Error("Could not disconnect from db", "error", err)
	}
}