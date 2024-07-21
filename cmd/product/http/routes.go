package http

import (
	"os"
	"product/cmd/product/handler"
	httphandler "product/internal/delivery/http"
	validatormiddleware "product/internal/delivery/http/middleware/validator"
	"product/internal/interactor"
)

// NewHandler initialize http service for the app
func NewHandler(interactor *interactor.AppInteractor) {
	router := httphandler.New()

	pHandler := &handler.ProductHandler{
		ProductService: interactor.ProductService,
	}

	cHandler := &handler.ChatHandler{
		ChatService: interactor.ChatService,
	}

	products := router.Group("/products")
	{
		products.GET("", pHandler.GetAll)
		products.GET("/:id", pHandler.GetProductByID)
		products.POST("", validatormiddleware.ValidateCreateProduct, pHandler.Create)
		products.PUT("/:id", validatormiddleware.ValidateCreateProduct, pHandler.Update)
		products.DELETE("/:id", pHandler.Delete)
	}
	router.POST("/ask", cHandler.AskAIHandler)

	router.Run(":" + os.Getenv("PORT"))
}
