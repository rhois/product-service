package product

import (
	"context"
	"product/internal/entity"
	"product/internal/presenter"
	"product/internal/repository"
)

// ServiceImpl struct to represent product service
type ServiceImpl struct {
	ProductRepo repository.Product
}

// Create product
func (s *ServiceImpl) Create(ctx context.Context, postData *presenter.Products) error {
	var product = entity.Products{
		Name:        postData.Name,
		Description: postData.Description,
		Price:       postData.Price,
		SupplierID:  postData.SupplierID,
	}
	err := s.ProductRepo.Create(ctx, &product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) Update(ctx context.Context, productID uint64, postData *presenter.Products) error {
	productData, err := s.ProductRepo.GetOne(ctx, productID)
	if err != nil {
		return err
	}

	updateData := entity.Products{
		ID:          productData.ID,
		Name:        postData.Name,
		Description: postData.Description,
		Price:       postData.Price,
		SupplierID:  postData.SupplierID,
	}
	err = s.ProductRepo.Update(ctx, &updateData)
	if err != nil {
		return err
	}
	return nil
}

// Delete selected product
func (s *ServiceImpl) Delete(ctx context.Context, productID uint64) error {
	res, err := s.ProductRepo.GetOne(ctx, productID)
	if err != nil {
		return err
	}
	err = s.ProductRepo.Delete(ctx, uint64(res.ID))
	if err != nil {
		return err
	}

	return nil
}

// GetAllProduct is function for get all product
func (s *ServiceImpl) GetAllProduct(ctx context.Context) ([]*entity.Products, error) {
	products, err := s.ProductRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID is function for get product by id
func (s *ServiceImpl) GetProductByID(ctx context.Context, productID uint64) (*entity.Products, error) {
	product, err := s.ProductRepo.GetOne(ctx, productID)
	if err != nil {
		return nil, err
	}
	return product, err
}
