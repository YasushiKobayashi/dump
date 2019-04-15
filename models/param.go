package models

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type (
	// Param
	Param struct {
		Output Output
		Type   ExportType

		Mysql     MysqlConnect
		SyncMysql MysqlConnect

		Path
		AwsS3 AwsS3Param
	}
)

func SetParam(c *cli.Context) (*Param, error) {
	output, err := ToOutput(c.String("output"))
	if err != nil {
		return nil, errors.Wrap(err, "unexpected output")
	}

	exportType, err := ToType(c.String("type"))
	if err != nil {
		return nil, errors.Wrap(err, "unexpected type")
	}

	res := &Param{
		Output: output,
		Type:   exportType,

		Mysql: MysqlConnect{
			User:     c.String("user"),
			Password: c.String("password"),
			Host:     c.String("host"),
			Port:     c.String("port"),
			Table:    c.String("table"),
		},
		SyncMysql: MysqlConnect{
			User:     c.String("sync_user"),
			Password: c.String("sync_password"),
			Host:     c.String("sync_host"),
			Port:     c.String("sync_port"),
			Table:    c.String("sync_table"),
		},

		Path: Path{
			FilePath: c.String("filepath"),
			FileName: c.String("filename"),
		},

		AwsS3: AwsS3Param{
			Bucket: c.String("bucket"),
			AwsId:  c.String("aws_id"),
			AwsKey: c.String("secret"),
			Region: c.String("region"),
		},
	}
	return res, nil
}

// Validate
func (c *Param) Validate() error {
	return nil
}
