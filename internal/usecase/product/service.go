package product

import (
	"context"
	"product/internal/entity"
	"product/internal/presenter"
)

// Service interface abstracts the controller layer and should be implemented in controller directory. Controller contains business logics and is independent of any database connection.
type Service interface {
	Create(ctx context.Context, postData *presenter.Products) error
	Update(ctx context.Context, productID uint64, postData *presenter.Products) error
	Delete(ctx context.Context, productID uint64) error
	GetAllProduct(ctx context.Context) ([]*entity.Products, error)
	GetProductByID(ctx context.Context, productID uint64) (*entity.Products, error)
}
