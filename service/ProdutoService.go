package service

import (
	"crud-go/entities"
	"crud-go/repository"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ProdutoService struct {
	repository *repository.ProdutoRepository
}

func NewProdutoService(repository *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repository,
	}
}

func validarProduto(produto entities.Produto) error {	//retorna um valor do tipo error
	if strings.TrimSpace(produto.Nome) == "" { 			// se o nome estiver vazio retorna erro
		return ErrNomeObrigatorio
	}

	if produto.Preco < 0 {  							//se preco for negativo retorna erro
		return ErrPrecoInvalido
	}

	return nil											//retorna um valor do tipo erro
}

func (s *ProdutoService) Save(produto *entities.Produto) error {
	if err := validarProduto(*produto); err != nil { 	//valida produto antes de salvar
		return err
	}

	return s.repository.Save(produto)
}

func (s *ProdutoService) FindAll() ([]entities.Produto, error) { 	//lista de produtos
	return s.repository.FindAll()
}

func (s *ProdutoService) FindByID(id int) (entities.Produto, error) {
	produto, err := s.repository.FindByID(id) 			//procura o produto pelo id, retorna produto e erro

	if errors.Is(err, gorm.ErrRecordNotFound) {				//O erro que aconteceu é do tipo gorm.ErrRecordNotFound?
		return entities.Produto{}, ErrProdutoNaoEncontrado	//erro personalizado
	}

	if err != nil {
		return entities.Produto{}, err
	}

	return produto, nil
}

func (s *ProdutoService) Delete(id int) error {
	err := s.repository.Delete(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {		//O erro que aconteceu é do tipo gorm.ErrRecordNotFound?
		return ErrProdutoNaoEncontrado
	}

	return err
}

func (s *ProdutoService) Update(id int, produtoAtualizado entities.Produto,) error {

	if err := validarProduto(produtoAtualizado); err != nil { 	//valida o produto atualizado
		return err
	}

	err := s.repository.Update(id, produtoAtualizado)

	if errors.Is(err, gorm.ErrRecordNotFound) {		//O erro que aconteceu é do tipo gorm.ErrRecordNotFound?
		return ErrProdutoNaoEncontrado
	}

	return err
}