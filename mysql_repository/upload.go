package mysql_repository

import (
	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

func (r *MysqlRepository) UploadDump(mysql *models.MysqlModel) error {
	mysql.Sort()
	for _, v := range mysql.SortedTables {
		query := mysql.CreateTables[v]
		_, err := r.DB.Exec(query)
		if err != nil {
			return errors.Wrapf(err, "r.DB.Exec error query is %s", query)
		}

		query = mysql.Inserts[v]
		if query != "" {
			_, err := r.DB.Exec(query)
			if err != nil {
				return errors.Wrap(err, "r.DB.Exec error")
			}
		}
	}
	return nil
}

func (r *MysqlRepository) UploadSync(mysql *models.MysqlModel) error {
	mysql.Sort()
	for _, v := range mysql.SortedTables {
		upserts := mysql.Upserts[v]
		for _, v := range upserts {
			_, err := r.DB.Exec(v)
			if err != nil {
				return errors.Wrapf(err, "r.DB.Exec error query is %s", v)
			}
		}
	}
	return nil
}
