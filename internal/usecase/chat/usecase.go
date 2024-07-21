package chat

import (
	"context"
	"product/internal/repository"
)

type ServiceImpl struct {
	OpenAIRepo repository.OpenAI
}

func (s *ServiceImpl) GenerateResponse(ctx context.Context, query string) (string, error) {
	res, err := s.OpenAIRepo.GenerateResponse(ctx, query)
	if err != nil {
		return "", err
	}
	return res, nil
}
