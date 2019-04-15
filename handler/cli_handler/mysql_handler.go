package cli_handler

import (
	"database/sql"

	"github.com/YasushiKobayashi/dump/aws_s3_repository"
	"github.com/YasushiKobayashi/dump/local_repository"
	"github.com/YasushiKobayashi/dump/models"
	"github.com/YasushiKobayashi/dump/mysql_repository"
	"github.com/YasushiKobayashi/dump/usecase"
	"github.com/urfave/cli"
)

var mysql = cli.Command{
	Name:  "mysql",
	Usage: "...",
	Description: `
...
`,
	Action: mysqlUploadLocal,
	Flags:  baseFlag,
}

func NewMysqlIdentifier(db *sql.DB, informationDB *sql.DB, targetDB string, syncDatabbase *sql.DB) *usecase.MysqlInterActor {
	return &usecase.MysqlInterActor{
		MysqlRepository: &mysql_repository.MysqlRepository{
			DB:            db,
			InformationDB: informationDB,
			TargetDB:      targetDB,
		},
		SyncMysqlRepository: &mysql_repository.MysqlRepository{
			DB: syncDatabbase,
		},
		LocalRepository: &local_repository.LocalRepository{},
		AwsS3Repository: &aws_s3_repository.AwsS3Repository{},
	}
}

func mysqlUploadLocal(c *cli.Context) error {
	param, err := models.SetParam(c)
	if err != nil {
		return cli.NewExitError(err, 128)
	}

	db := mysql_repository.ConnectDB(param.Mysql)
	informationDB := mysql_repository.ConnectInformationSchema(param.Mysql)
	defer func() {
		db.Close()
		informationDB.Close()
	}()

	var syncDatabbase *sql.DB = nil
	if param.Output == models.OtherDatabase {
		syncDatabbase = mysql_repository.ConnectDB(param.SyncMysql)
		defer func() {
			syncDatabbase.Close()
		}()
	}

	h := NewMysqlIdentifier(db, informationDB, param.Mysql.Table, syncDatabbase)
	err = h.Upload(param)
	if err != nil {
		return cli.NewExitError(err, 128)
	}
	return nil
}
