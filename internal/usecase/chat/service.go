package chat

import "context"

type Service interface {
	GenerateResponse(ctx context.Context, query string) (string, error)
}
