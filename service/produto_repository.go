package service

import "crud-go/entities"

type ProdutoRepositoryInterface interface {
	Save(produto entities.Produto) error
	FindAll() ([]entities.Produto, error)
	FindByID(id int) (entities.Produto, error)
	Delete(id int) error
	Update(id int, produto entities.Produto) error
}