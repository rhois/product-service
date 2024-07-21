package repository

import (
	"context"
	"product/internal/entity"
)

// Product interface abstracts the repository layer and should be implemented in repository
// directory. Filename should be prefixed using the data retrieval method such as database,
// or microservice connection. example : `psqlrepository.go`
type Product interface {
	Create(ctx context.Context, data *entity.Products) error
	Update(ctx context.Context, data *entity.Products) error
	GetOne(ctx context.Context, ID uint64) (*entity.Products, error)
	GetAll(ctx context.Context) ([]*entity.Products, error)
	Delete(ctx context.Context, productID uint64) error
}
