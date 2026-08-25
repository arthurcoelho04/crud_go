package repository

import "crud-go/entities"
import "errors"

type ProdutoRepository struct {
	produtos []entities.Produto
}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{
		produtos: []entities.Produto{},
	}
}

func (r *ProdutoRepository) Save(produto entities.Produto) {
	r.produtos = append(r.produtos, produto)
}

func (r *ProdutoRepository) FindAll() []entities.Produto { //retorna uma lista de produtos
	return r.produtos
}

func (r *ProdutoRepository) FindByID(id int) (entities.Produto, error) {

    for _, produto := range r.produtos {

        if produto.ID == id {
            return produto, nil
        }
    }

    return entities.Produto{}, errors.New("produto não encontrado")
}
