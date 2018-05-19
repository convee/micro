package model

import "gitlab.com/convee/system/mysql"

type Model struct {
	db *mysql.Mysql
}

func NewModel() *Model {
	name := "chess"
	return &Model{db: mysql.New(name)}
}
