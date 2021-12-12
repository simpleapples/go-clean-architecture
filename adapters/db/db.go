package db

import "clean_architecutre_demo/entities"

type IDBAdapter interface {
	QueryArticleByID(id int) *entities.Article
}

func NewDBAdapater() IDBAdapter {
	return newMysqlDBAdapter()
}
