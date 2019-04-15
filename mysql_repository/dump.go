package mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

func (r *MysqlRepository) GetDump() (*models.MysqlModel, error) {
	res := &models.MysqlModel{}
	err := r.getTableNames(res)
	if err != nil {
		return res, errors.Wrap(err, "r.getTableNames error")
	}

	err = r.getRelation(res)
	if err != nil {
		return res, errors.Wrap(err, "r.getRelation error")
	}

	err = r.getTableInfo(res)
	if err != nil {
		return res, errors.Wrap(err, "r.getTableInfo error")
	}

	err = r.getInsertInfo(res)
	if err != nil {
		return res, errors.Wrap(err, "r.getInsertInfo error")
	}

	return res, nil
}

func (r *MysqlRepository) getInsertInfo(d *models.MysqlModel) error {
	d.Inserts = make(map[models.Table]string, len(d.SortedTables))
	for _, v := range d.Tables {
		insert, err := r.getInsert(string(v))
		if err != nil {
			return errors.Wrap(err, "r.getCreateTable error")
		}
		d.Inserts[v] = insert
	}
	return nil
}

func (r *MysqlRepository) getInsert(tableName string) (res string, err error) {
	rows, err := r.DB.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return res, errors.Wrap(err, `"SELECT * FROM ?" error`)
	}
	defer rows.Close()

	count, err := r.countRecord(tableName)
	if err != nil {
		return res, errors.Wrap(err, "r.countRecord error")
	}

	// Get columns
	columns, err := rows.Columns()
	if err != nil {
		return res, errors.Wrap(err, "rows.Columns error")
	}

	columnSizes := len(columns)
	if columnSizes == 0 {
		return res, errors.Errorf("No columns in table %s", tableName)
	}

	insertText := make([]string, len(columns))
	for k, v := range columns {
		insertText[k] = fmt.Sprintf(`%s`, v)
	}

	dataText := make([]string, count)
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

		dataStrings := make([]string, len(columns))
		for key, value := range data {
			if value != nil && value.Valid {
				dataStrings[key] = fmt.Sprintf(`'%s'`, value.String)
			} else {
				dataStrings[key] = null
			}
		}

		dataText[i] = fmt.Sprintf("(%s)", strings.Join(dataStrings, ", "))
		i = i + 1
	}

	err = rows.Err()
	if err != nil {
		return res, errors.Wrap(err, "rows.Err")
	}

	if len(dataText) == 0 {
		return "", nil
	}

	script := fmt.Sprintf(`INSERT INTO %s `+"(`%s`)"+` VALUES
%s;
`, tableName, strings.Join(insertText, "`, `"), strings.Join(dataText, ", "+`
`))
	return script, nil
}
