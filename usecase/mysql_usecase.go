package usecase

import (
	"fmt"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

type (
	MysqlInterActor struct {
		MysqlRepository     MysqlRepository
		SyncMysqlRepository MysqlRepository
		LocalRepository     LocalRepository
		AwsS3Repository     AwsS3Repository
	}

	MysqlRepository interface {
		GetDump() (*models.MysqlModel, error)
		GetSync() (*models.MysqlModel, error)
		UploadDump(*models.MysqlModel) error
		UploadSync(*models.MysqlModel) error
	}

	LocalRepository interface {
		Upload(string, models.Path) error
	}

	AwsS3Repository interface {
		Upload(string, models.Path, models.AwsS3Param) error
	}
)

func (i *MysqlInterActor) Upload(param *models.Param) (err error) {
	var mysql *models.MysqlModel
	if param.Type == models.Dump {
		mysql, err = i.MysqlRepository.GetDump()
		if err != nil {
			return errors.Wrap(err, "i.MysqlRepository.GetDump error")
		}
	}

	if param.Type == models.Sync {
		mysql, err = i.MysqlRepository.GetSync()
		if err != nil {
			return errors.Wrap(err, "i.MysqlRepository.GetSync error")
		}
	}

	if param.Output == models.Stout {
		query := mysql.GetQuery(param.Type)
		fmt.Println(query)
		return nil
	}

	if param.Output == models.File {
		query := mysql.GetQuery(param.Type)
		err := i.LocalRepository.Upload(query, param.Path)
		if err != nil {
			return errors.Wrap(err, "i.LocalRepository.Upload error")
		}
		return nil
	}

	if param.Output == models.AwsS3 {
		query := mysql.GetQuery(param.Type)
		err := i.AwsS3Repository.Upload(query, param.Path, param.AwsS3)
		if err != nil {
			return errors.Wrap(err, "i.AwsS3Repository.Upload error")
		}
		return nil
	}

	if param.Output == models.OtherDatabase {
		if param.Type == models.Dump {
			err := i.SyncMysqlRepository.UploadDump(mysql)
			if err != nil {
				return errors.Wrap(err, "i.SyncMysqlRepository.UploadDump error")
			}
			return nil
		}

		if param.Type == models.Sync {
			err := i.SyncMysqlRepository.UploadSync(mysql)
			if err != nil {
				return errors.Wrap(err, "i.SyncMysqlRepository.UploadSync error")
			}
			return nil
		}
	}
	return nil
}
