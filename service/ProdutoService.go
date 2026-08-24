package service

import (
	"crud-go/entities"
	"crud-go/repository"
)

type ProdutoService struct {
	repository *repository.ProdutoRepository
}

func NewProdutoService(repository *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repository,
	}
}

func (s *ProdutoService) Save(produto entities.Produto) { //metodo salvar produto
	s.repository.Save(produto)
}

func (s *ProdutoService) FindAll() []entities.Produto {
	return s.repository.FindAll()
}
