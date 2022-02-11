package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/verryp/cake-store-service/app/common"
	"gopkg.in/gorp.v2"
)

func NewMySQLDatabase(opt common.DB) (*gorp.DbMap, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	return dbMap, nil
}
