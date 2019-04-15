package mysql_repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func Test_getTableCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tableName := "test_table"

	countRes := 2

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(countRes)
	mock.ExpectQuery("^SELECT (.+) FROM ?").WillReturnRows(rows)

	repository := &MysqlRepository{
		DB: db,
	}

	count, err := repository.countRecord(tableName)
	if err != nil {
		t.Fatalf("faile %#v", err)
	}
	if count != countRes {
		t.Fatalf("faile test, expected %d, but get %#v", countRes, count)
	}
}
