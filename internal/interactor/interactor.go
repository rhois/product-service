package interactor

import (
	"product/internal/usecase/chat"
	"product/internal/usecase/product"
)

// AppInteractor represents app's interactor object
type AppInteractor struct {
	ProductService product.Service
	ChatService    chat.Service
}
