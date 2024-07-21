package psql

import (
	"context"
	"product/internal/entity"
	"product/internal/repository"
	"time"

	"github.com/jinzhu/gorm"
)

type psqlProductRepository struct {
	Conn *gorm.DB
}

// NewProductRepository function is used to initialize repository
// implementing the functions defined in product's
// repository interface
func NewProductRepository(conn *gorm.DB) repository.Product {
	return &psqlProductRepository{
		Conn: conn,
	}
}

func (p *psqlProductRepository) Create(ctx context.Context, data *entity.Products) error {
	return p.Conn.Create(data).Error
}

func (p *psqlProductRepository) Update(ctx context.Context, data *entity.Products) error {
	return p.Conn.Save(data).Error
}

func (p *psqlProductRepository) GetOne(ctx context.Context, ID uint64) (*entity.Products, error) {
	var result entity.Products
	err := p.Conn.Where("id = ?", ID).Find(&result).Error
	return &result, err
}

func (p *psqlProductRepository) GetAll(ctx context.Context) ([]*entity.Products, error) {
	var result []*entity.Products

	err := p.Conn.Find(&result).Error
	return result, err
}

func (p *psqlProductRepository) Delete(ctx context.Context, productID uint64) error {
	timeNow := time.Now()
	var product entity.Products
	return p.Conn.Model(product).Where("id = ?", productID).Update(entity.Products{
		DeletedAt: &timeNow,
	}).Error
}
