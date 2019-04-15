package cli_handler

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	mysql,
}

func margeBaseFlag(f []cli.Flag) []cli.Flag {
	return append(baseFlag, f...)
}

var baseFlag []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:   "output, o",
		Value:  "stdout",
		EnvVar: "output",
		Usage:  "stdout, file, aws-s3, other-database",
	},
	cli.StringFlag{
		Name:   "type, t",
		Value:  "dump",
		EnvVar: "TYPE",
		Usage:  "dump, sync",
	},

	cli.StringFlag{
		Name:   "user, u",
		Value:  "root",
		EnvVar: "DB_USER",
		Usage:  "database user",
	},
	cli.StringFlag{
		Name:   "password, p",
		Value:  "",
		EnvVar: "DB_PASSWORD",
		Usage:  "database password",
	},
	cli.StringFlag{
		Name:   "host",
		Value:  "127.0.0.1",
		EnvVar: "DB_HOST",
		Usage:  "database host",
	},
	cli.StringFlag{
		Name:   "port",
		Value:  "3306",
		EnvVar: "DB_PORT",
		Usage:  "database port",
	},
	cli.StringFlag{
		Name:   "table",
		Value:  "",
		EnvVar: "DB_TABLE",
		Usage:  "database table",
	},
	cli.StringFlag{
		Name:   "sync_user, su",
		Value:  "root",
		EnvVar: "SYNC_DB_USER",
		Usage:  "sync database user",
	},
	cli.StringFlag{
		Name:   "sync_password, sp",
		Value:  "",
		EnvVar: "SYNC_DB_PASSWORD",
		Usage:  "sync database password",
	},
	cli.StringFlag{
		Name:   "sync_host, sh",
		Value:  "127.0.0.1",
		EnvVar: "SYNC_DB_HOST",
		Usage:  "sync database host",
	},
	cli.StringFlag{
		Name:   "sync_port",
		Value:  "3306",
		EnvVar: "SYNC_DB_PORT",
		Usage:  "sync database port",
	},
	cli.StringFlag{
		Name:   "sync_table, st",
		Value:  "",
		EnvVar: "SYNC_DB_TABLE",
		Usage:  "sync database table",
	},

	cli.StringFlag{
		Name:   "filepath",
		Value:  "./",
		EnvVar: "FILE_PATH",
		Usage:  "dump set file path",
	},
	cli.StringFlag{
		Name:   "filename, n",
		Value:  fmt.Sprintf("dump-%s", time.Now().Format("20060102")),
		EnvVar: "FILE_NAME",
		Usage:  "filename",
	},

	cli.StringFlag{
		Name:   "bucket, b",
		Value:  "",
		EnvVar: "S3_BUCKET",
		Usage:  "s3 bucket",
	},
	cli.StringFlag{
		Name:   "aws_id, i",
		Value:  "",
		EnvVar: "AWS_ACCESS_KEY_ID",
		Usage:  "aws iam user id",
	},
	cli.StringFlag{
		Name:   "secret, s",
		Value:  "",
		EnvVar: "AWS_SECRET_ACCESS_KEY",
		Usage:  "aws iam user secret access key",
	},
	cli.StringFlag{
		Name:   "region, r",
		Value:  "us-west-1",
		EnvVar: "AWS_REGION",
		Usage:  "aws region",
	},
}
