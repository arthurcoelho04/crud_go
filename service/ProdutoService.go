package service

import (
	"crud-go/entities"
	"strings"
)

type ProdutoService struct {
	repository ProdutoRepositoryInterface
}

func NewProdutoService(repository ProdutoRepositoryInterface) *ProdutoService {
	return &ProdutoService{
		repository: repository,
	}
}

func validarProduto(produto entities.Produto) error {
	if strings.TrimSpace(produto.Nome) == "" {
		return ErrNomeObrigatorio
	}

	if produto.Preco < 0 {
		return ErrPrecoInvalido
	}

	return nil
}

func (s *ProdutoService) Save(produto entities.Produto) error {
	if err := validarProduto(produto); err != nil {
		return err
	}

	return s.repository.Save(produto)
}

func (s *ProdutoService) FindAll() ([]entities.Produto, error) {
	return s.repository.FindAll()
}

func (s *ProdutoService) FindByID(id int) (entities.Produto, error) {
	return s.repository.FindByID(id)
}

func (s *ProdutoService) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *ProdutoService) Update(id int, produtoAtualizado entities.Produto) error {
	if err := validarProduto(produtoAtualizado); err != nil {
		return err
	}

	return s.repository.Update(id, produtoAtualizado)
}