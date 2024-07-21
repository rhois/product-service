package repository

import "context"

type OpenAI interface {
	GenerateResponse(ctx context.Context, query string) (string, error)
}
