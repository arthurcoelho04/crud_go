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

// obs// Manda o produto para o Repository salvar no banco
func (s *ProdutoService) Save(produto entities.Produto) error {
	return s.repository.Save(produto)
}

// obs// Busca todos os produtos através do Repository
func (s *ProdutoService) FindAll() ([]entities.Produto, error) {
	return s.repository.FindAll()
}

// obs// Busca um produto pelo ID através do Repository
func (s *ProdutoService) FindByID(id int) (entities.Produto, error) {
	return s.repository.FindByID(id)
}

// obs// Manda o ID para o Repository deletar o produto
func (s *ProdutoService) Delete(id int) error {
	return s.repository.Delete(id)
}

// obs// Manda os dados atualizados para o Repository
func (s *ProdutoService) Update(id int, produtoAtualizado entities.Produto) error {
	return s.repository.Update(id, produtoAtualizado)
}