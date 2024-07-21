package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"product/internal/presenter"
	"product/internal/repository"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/jinzhu/gorm"
	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/api/option"
)

type openAIRepository struct {
	Conn         *gorm.DB
	OpenAIClient string
	GeminiKey    string
}

// NewOpenaiRepository function is used to initialize repository
// implementing the functions defined in openai
// repository interface
func NewOpenaiRepository(conn *gorm.DB, client, apikey string) repository.OpenAI {
	return &openAIRepository{
		Conn:         conn,
		OpenAIClient: client,
		GeminiKey:    apikey,
	}
}

func (r *openAIRepository) GenerateResponse(ctx context.Context, query string) (string, error) {

	var res string

	result, err := r.GetProductWithSupplier(ctx)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	s := string(b)
	s = strings.TrimSpace(s)
	// trim leading and trailing spaces
	s = s[1 : len(s)-1]
	res = s

	prompt := fmt.Sprintf("Based on the following data:\n%s\nAnswer the query: %s", res, query)
	if r.GeminiKey != "" {
		client, err := genai.NewClient(ctx, option.WithAPIKey(r.GeminiKey))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
		model := client.GenerativeModel("gemini-1.5-flash")
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			log.Fatal(err)
		}
		res = printResponse(resp)
	} else if r.OpenAIClient != "" {
		client := openai.NewClient(r.OpenAIClient)
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 20,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		}
		response, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			return "", err
		}
		res = response.Choices[0].Message.Content
	}

	return res, nil
}

func printResponse(resp *genai.GenerateContentResponse) string {
	var res string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				res = fmt.Sprintf("%s", part)
			}
		}
	}
	return res
}

func (r *openAIRepository) GetProductWithSupplier(ctx context.Context) ([]presenter.ProductsWithSuppliers, error) {
	var productsWithSuppliers []presenter.ProductsWithSuppliers

	// Define the SQL query for joining products and suppliers
	query := `
        SELECT
            products.id AS product_id,
            products.name AS product_name,
            products.description,
            products.price,
            suppliers.id AS supplier_id,
            suppliers.name AS supplier_name,
            suppliers.contact_info
        FROM
            products
        JOIN
            suppliers
        ON
            products.supplier_id = suppliers.id;
    `

	// Execute the query
	if err := r.Conn.Raw(query).Scan(&productsWithSuppliers).Error; err != nil {
		return nil, err
	}

	return productsWithSuppliers, nil
}
