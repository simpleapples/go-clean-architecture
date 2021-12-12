package usecases

import (
	"clean_architecutre_demo/adapters/db"
	"clean_architecutre_demo/entities"
)

type articleUsecase struct {
	dbAdapter db.IDBAdapter
}

func NewArticleUsecase() IArticleUsecase {
	dbAdapter := db.NewDBAdapater()
	return &articleUsecase{dbAdapter: dbAdapter}
}

func (u *articleUsecase) Show(id int) *entities.Article {
	return u.dbAdapter.QueryArticleByID(id)
}
