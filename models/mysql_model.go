package models

import (
	"strings"
)

type (
	MysqlModel struct {
		Tables       Tables
		Tbs          Tables
		SortedTables Tables
		References   References
		CreateTables map[Table]string
		Inserts      map[Table]string
		Upserts      map[Table][]string
	}

	Table string

	Tables     []Table
	References []Reference

	Reference struct {
		TableName           Table `db:"TABLE_NAME"`
		ReferencedTableName Table `db:"REFERENCED_TABLE_NAME"`
	}
)

func (d *MysqlModel) GetQuery(t ExportType) string {
	d.Sort()
	var res []string
	if t == Dump {
		res = []string{}
		for _, v := range d.SortedTables {
			res = append(res, d.CreateTables[v])
			res = append(res, d.Inserts[v])
		}
	}

	if t == Sync {
		res = []string{}
		for _, v := range d.SortedTables {
			upserts := d.Upserts[v]
			for _, v := range upserts {
				res = append(res, v)
			}
		}
	}
	return strings.Join(res, "\n")
}

func (d *MysqlModel) Sort() {
	tableLength := len(d.Tables)
	d.Tbs = d.Tables
	for {
		for _, tb := range d.Tbs {
			d.forParents(tb)
		}

		if tableLength == len(d.SortedTables) {
			break
		}
	}
}

func (d *MysqlModel) forParents(tb Table) {
	parents := d.getParent(tb)
	if len(parents) == 0 || d.existsSorted(parents) {
		d.addSorted(tb)
	} else {
		for _, v := range parents {
			grandParents := d.getParent(v)
			if len(grandParents) == 0 || d.existsSorted(grandParents) {
				d.addSorted(v)
			} else {
				if !d.SortedTables.exists(v) {
					d.forParents(v)
				}
			}
		}
	}
}

func (d *MysqlModel) addSorted(table Table) {
	if !d.SortedTables.exists(table) {
		d.SortedTables = append(d.SortedTables, table)
		d.Tbs = d.Tbs.removeTable(table)
	}
}

func (r Tables) exists(tableName Table) bool {
	for _, v := range r {
		if v == tableName {
			return true
		}
	}
	return false
}

func (d Tables) removeTable(tableName Table) Tables {
	var res Tables
	for _, v := range d {
		if v != tableName {
			res = append(res, v)
		}
	}
	return res
}

func (d *MysqlModel) getParent(tableName Table) []Table {
	var res []Table
	for _, v := range d.References {
		if v.TableName == tableName {
			res = append(res, d.Tables.getTable(v.ReferencedTableName))
		}
	}
	return res
}

func (d *MysqlModel) existsSorted(tableNames Tables) bool {
	count := 0
	for _, tableName := range tableNames {
		for _, sortedTable := range d.SortedTables {
			if tableName == sortedTable {
				count = count + 1
				if count == len(tableNames) {
					return true
				}
			}
		}
	}
	return false
}

func (t Tables) getTable(tableName Table) Table {
	var res Table
	for _, v := range t {
		if v == tableName {
			res = v
		}
	}
	return res
}
