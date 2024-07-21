package main

import (
	"os"

	"product/cmd/product/http"
	dbinfra "product/internal/infrastructure/db/psql"
	"product/internal/interactor"
	openaiRepo "product/internal/repository/openai"
	psqlrepo "product/internal/repository/psql"
	"product/internal/usecase/chat"
	"product/internal/usecase/product"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// Initialize infrastuctures
	dbURL := os.Getenv("DATABASE_URL")
	openAIKey := os.Getenv("OPENAI_API_KEY")
	geminiKey := os.Getenv("GEMINI_API_KEY")

	// Psql database
	psqlRepo := dbinfra.NewPsqlRepository(dbURL)
	defer psqlRepo.Close()

	// List all services required by the application
	// psql repository
	productRepo := psqlrepo.NewProductRepository(psqlRepo)
	chatRepo := openaiRepo.NewOpenaiRepository(psqlRepo, openAIKey, geminiKey)

	// product service
	productSvc := &product.ServiceImpl{
		ProductRepo: productRepo,
	}

	// chat service
	chatSvc := &chat.ServiceImpl{
		OpenAIRepo: chatRepo,
	}

	// Tidying up all services to an interactor
	interactor := &interactor.AppInteractor{
		ProductService: productSvc,
		ChatService:    chatSvc,
	}

	http.NewHandler(interactor)
}
