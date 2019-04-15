package mysql_repository

import (
	"database/sql"
	"fmt"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/pkg/errors"
)

type (
	MysqlRepository struct {
		DB            *sql.DB
		InformationDB *sql.DB
		TargetDB      string
	}
)

const null = "NULL"

func (r *MysqlRepository) countRecord(tableName string) (int, error) {
	var count int
	err := r.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		return count, errors.Wrap(err, `"SELECT COUNT(*) FROM ?" error`)
	}
	return count, nil
}

func (r *MysqlRepository) getTableInfo(d *models.MysqlModel) error {
	d.CreateTables = make(map[models.Table]string)
	for _, v := range d.Tables {
		createTable, err := r.getCreateTable(string(v))
		if err != nil {
			return errors.Wrap(err, "r.getCreateTable error")
		}
		d.CreateTables[v] = createTable
	}
	return nil
}

func (r *MysqlRepository) getCreateTable(tableName string) (res string, err error) {
	var tableReturn sql.NullString
	var tableSql sql.NullString

	err = r.DB.QueryRow(fmt.Sprintf("SHOW CREATE TABLE %s;", tableName)).Scan(&tableReturn, &tableSql)
	if err != nil {
		return res, errors.Wrap(err, "SHOW CREATE TABLE err")
	}

	if tableReturn.String != tableName {
		return res, errors.New("Returned table is not the same as requested table")
	}

	res = fmt.Sprintf(`%s;`, tableSql.String)
	return res, nil
}

func (r *MysqlRepository) getTableNames(d *models.MysqlModel) error {
	rows, err := r.DB.Query("SHOW TABLES;")
	if err != nil {
		return errors.Wrap(err, "SHOW TABLES query error")
	}
	defer rows.Close()

	for rows.Next() {
		var table sql.NullString
		if err := rows.Scan(&table); err != nil {
			return errors.Wrap(err, fmt.Sprintf("rows.Scan %s error", table.String))
		}
		d.Tables = append(d.Tables, models.Table(table.String))
	}
	return nil
}

func (r *MysqlRepository) getRelation(d *models.MysqlModel) error {
	rows, err := r.InformationDB.
		Query(`SELECT TABLE_NAME, REFERENCED_TABLE_NAME FROM REFERENTIAL_CONSTRAINTS WHERE CONSTRAINT_SCHEMA = ?`, r.TargetDB)
	if err != nil {
		return errors.Wrap(err, `"SELECT TABLE_NAME, REFERENCED_TABLE_NAME FROM REFERENTIAL_CONSTRAINTS WHERE CONSTRAINT_SCHEMA = ?" error`)
	}

	for rows.Next() {
		var reference models.Reference
		err := rows.Scan(&reference.TableName, &reference.ReferencedTableName)
		if err != nil {
			return errors.Wrap(err, "scan error")
		}
		d.References = append(d.References, reference)
	}

	return nil
}
