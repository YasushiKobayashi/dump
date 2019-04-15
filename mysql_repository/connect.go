package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/YasushiKobayashi/dump/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func dbHostUrl(dbUser, dbPassword, dbHost, dbPort, dbTable string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbTable)
}

func ConnectDB(flags models.MysqlConnect) *sql.DB {
	db, err := sql.Open("mysql", dbHostUrl(flags.User, flags.Password, flags.Host, flags.Port, flags.Table))
	if err != nil {
		panic(errors.Wrap(err, "connect error"))
	}
	return db
}

func ConnectInformationSchema(flags models.MysqlConnect) *sql.DB {
	db, err := sql.Open("mysql", dbHostUrl(flags.User, flags.Password, flags.Host, flags.Port, "information_schema"))
	if err != nil {
		panic(errors.Wrap(err, "connect error"))
	}
	return db
}
