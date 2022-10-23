package services

import (
	"context"

	"github.com/kingofmidas/astrology-api/internal/pkg/dto"
	"github.com/kingofmidas/astrology-api/internal/pkg/entity"
)

type ImageRepository interface {
	GetAll(ctx context.Context) ([]entity.Image, error)
	GetByDate(ctx context.Context, date string) (entity.Image, error)
	Create(ctx context.Context, image entity.Image) (entity.Image, error)
}

type CollectorService interface {
	GetAPOD(ctx context.Context, params dto.ApodQueryParams) (entity.Image, error)
}
