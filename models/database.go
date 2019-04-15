package models

type (
	MysqlConnect struct {
		User     string
		Password string
		Host     string
		Port     string
		Table    string
	}
)
