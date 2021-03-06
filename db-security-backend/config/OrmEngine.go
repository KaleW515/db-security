package config

import (
	"db-security-backend/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.CharSet
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return nil, err
	}
	// engine.ShowSQL(database.ShowSql)
	err = engine.Sync2(
		new(model.User), new(model.BannedIp), new(model.Patient), new(model.Config), new(model.Staff), new(model.DownloadRecord),
	)
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, err
}
