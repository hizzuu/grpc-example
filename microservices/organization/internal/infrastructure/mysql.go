package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var dirver string = "mysql"

func NewMysqlDB() (*sql.DB, error) {
	c := &mysql.Config{
		User:                 DBConf.User,
		Passwd:               DBConf.Pass,
		Net:                  DBConf.Net,
		Addr:                 DBConf.Host + ":" + DBConf.Port,
		DBName:               DBConf.Name,
		Loc:                  time.Local,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open(dirver, c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
