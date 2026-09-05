package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) *ProdutoRepository {
	return &ProdutoRepository{
		db: db,
	}
}

// obs// Salva um novo produto no banco de dados
func (r *ProdutoRepository) Save(produto entities.Produto) error {

	err := r.db.Create(&produto).Error

	return err
}

// obs// Get
func (r *ProdutoRepository) FindAll() ([]entities.Produto, error) {
	var produtos []entities.Produto

	err := r.db.Find(&produtos).Error

	return produtos, err
}

// obs// GetByID
func (r *ProdutoRepository) FindByID(id int) (entities.Produto, error) {
	var produto entities.Produto

	err := r.db.First(&produto, id).Error

	return produto, err
}

// obs// DeleteByID
func (r *ProdutoRepository) Delete(id int) error {
	return r.db.Delete(&entities.Produto{}, id).Error
}

// obs// Atualiza um produto pelo ID
func (r *ProdutoRepository) Update(id int, produtoAtualizado entities.Produto) error {
	return r.db.Model(&entities.Produto{}).
		Where("id = ?", id).
		Updates(produtoAtualizado).Error
}
