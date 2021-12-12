package db

import "clean_architecutre_demo/entities"

type mysqlDBAdapter struct {
}

func newMysqlDBAdapter() IDBAdapter {
	return &mysqlDBAdapter{}
}

func (a *mysqlDBAdapter) QueryArticleByID(id int) *entities.Article {
	// you may call gorm to query an Article from MySQL here.
	// just build and return an Article instance instead of deploying MySQL.
	articleFromDB := entities.Article{
		ID:      1,
		Title:   "article title",
		Content: "article content",
	}
	return &articleFromDB
}
