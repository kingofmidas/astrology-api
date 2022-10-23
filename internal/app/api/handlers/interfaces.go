package handlers

import (
	"context"

	"github.com/kingofmidas/astrology-api/internal/pkg/entity"
)

type ImageService interface {
	GetAll(ctx context.Context) ([]entity.Image, error)
	GetByDate(ctx context.Context, date string) (entity.Image, error)
}
