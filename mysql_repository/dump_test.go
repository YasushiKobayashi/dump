package mysql_repository

import (
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func Test_getInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tableName := "test_table"

	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "body", "deleted"}).
		AddRow(1, 1, "post 1", `{json": ["rigid"]}`, nil).
		AddRow(2, 2, "post 1", `{json": ["rigid"]}`, nil)
	mock.ExpectQuery("^SELECT (.+) FROM ?").WillReturnRows(rows)

	countRes := 2
	rows = sqlmock.NewRows([]string{"count"}).
		AddRow(countRes)
	mock.ExpectQuery("^SELECT (.+) FROM ?").WillReturnRows(rows)

	repository := &MysqlRepository{
		DB: db,
	}

	insert, err := repository.getInsert(tableName)
	if err != nil {
		t.Fatalf("faile %#v", err)
	}
	if !strings.Contains(insert, "(`id`, `user_id`, `title`, `body`, `deleted`)") ||
		!strings.Contains(insert, `('1', '1', 'post 1', '{json": ["rigid"]}', NULL)`) {
		t.Fatalf("failed test %#v", insert)
	}
}
