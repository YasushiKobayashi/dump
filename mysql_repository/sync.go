package mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

func (r *MysqlRepository) GetSync() (*models.MysqlModel, error) {
	res := &models.MysqlModel{}

	err := r.getTableNames(res)
	if err != nil {
		return res, errors.Wrap(err, "GetTableNames error")
	}

	err = r.getRelation(res)
	if err != nil {
		return res, errors.Wrap(err, "GetTableNames error")
	}

	err = r.getTableInfo(res)
	if err != nil {
		return res, errors.Wrap(err, "getTableInfo error")
	}

	err = r.getUpsertInfo(res)
	if err != nil {
		return res, errors.Wrap(err, "getUpsertInfo error")
	}
	return res, nil
}

func (r *MysqlRepository) getUpsertInfo(d *models.MysqlModel) error {
	d.Upserts = make(map[models.Table][]string, len(d.Tables))
	for _, v := range d.Tables {
		upsert, err := r.getUpsert(string(v))
		if err != nil {
			return errors.Wrap(err, "r.getUpsert error")
		}
		d.Upserts[v] = upsert
	}
	return nil
}

func (r *MysqlRepository) getUpsert(tableName string) (res []string, err error) {
	rows, err := r.DB.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return res, errors.Wrap(err, `"SELECT * FROM ?" error`)
	}
	defer rows.Close()

	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return res, errors.Wrap(err, "rows.Columns error")
	}

	columnSizes := len(columns)
	if columnSizes == 0 {
		return res, errors.Errorf("No columns in table %s", tableName)
	}

	insertTexts := make([]string, len(columns))
	for k, v := range columns {
		insertTexts[k] = fmt.Sprintf(`%s`, v)
	}

	count, err := r.countRecord(tableName)
	if err != nil {
		return res, errors.Wrap(err, "r.countRecord error")
	}

	upserts := make([][]string, count)
	dataText := make([][]string, count)
	i := 0
	for rows.Next() {
		data := make([]*sql.NullString, len(columns))
		ptrs := make([]interface{}, len(columns))
		for i, _ := range data {
			ptrs[i] = &data[i]
		}

		// Read data
		if err := rows.Scan(ptrs...); err != nil {
			return res, errors.Wrap(err, "rows.Scan error")
		}

		upsert := make([]string, len(columns))
		dataStrings := make([]string, len(columns))
		for key, value := range data {
			if value != nil && value.Valid {
				dataStrings[key] = fmt.Sprintf(`'%s'`, value.String)
				upsert[key] = fmt.Sprintf("`%s`"+`='%s'`, columns[key], value.String)
			} else {
				dataStrings[key] = null
				upsert[key] = fmt.Sprintf("`%s`"+`=%s`, columns[key], null)
			}
		}

		upserts[i] = upsert
		dataText[i] = dataStrings
		i = i + 1
	}

	err = rows.Err()
	if err != nil {
		return res, errors.Wrap(err, "rows.Err")
	}

	if len(dataText) == 0 {
		return res, nil
	}

	insertText := strings.Join(insertTexts, "`, `")

	scripts := make([]string, len(upserts))
	for k, v := range upserts {
		script := fmt.Sprintf(`INSERT INTO %s `+"(`%s`)"+` VALUES
`+"(%s)"+`
ON DUPLICATE KEY UPDATE
%s;
`, tableName, insertText, strings.Join(dataText[k], ", "), strings.Join(v, ", \n"))
		scripts[k] = script

	}
	return scripts, nil
}
