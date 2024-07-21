package handler

import (
	"net/http"
	"product/internal/presenter"
	"product/internal/usecase/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler holds app interactor to be used by handler function
type ProductHandler struct {
	ProductService product.Service
}

// Create is function for create product
func (h *ProductHandler) Create(c *gin.Context) {
	data := c.MustGet("post_data").(presenter.Products)
	err := h.ProductService.Create(c, &data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"ok":      false,
			"code":    "unprocessable-entity",
			"errors":  []string{err.Error()},
			"message": "Error when trying to create product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Successfully create product",
	})

}

// Update for update product
func (h *ProductHandler) Update(c *gin.Context) {
	data := c.MustGet("post_data").(presenter.Products)
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := h.ProductService.Update(c, productID, &data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"ok":      false,
			"code":    "unprocessable-entity",
			"errors":  []string{err.Error()},
			"message": "Error when trying to update product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Successfully update product",
	})

}

// Delete for delete product
func (h *ProductHandler) Delete(c *gin.Context) {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := h.ProductService.Delete(c, productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"ok":      false,
			"code":    "unprocessable-entity",
			"errors":  []string{err.Error()},
			"message": "Error when trying to delete product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Successfully delete product",
	})

}

// GetAll for get all product
func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.ProductService.GetAllProduct(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"code":    "internal-server-error",
			"errors":  []string{},
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":       true,
		"message":  "Succesfully get products",
		"products": products,
	})
}

// GetByID for get product by id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	product, err := h.ProductService.GetProductByID(c, productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"code":    "internal-server-error",
			"errors":  []string{},
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "succesfully get detail product",
		"product": product,
	})
}
