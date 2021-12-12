package usecases

import "clean_architecutre_demo/entities"

type IArticleUsecase interface {
	Show(id int) *entities.Article
}
