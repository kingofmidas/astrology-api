package services

import (
	"context"
	"fmt"

	"github.com/kingofmidas/astrology-api/internal/pkg/dto"
	"github.com/kingofmidas/astrology-api/internal/pkg/entity"
	errs "github.com/kingofmidas/astrology-api/internal/pkg/errors"
)

type imageService struct {
	imageRepository  ImageRepository
	collectorService CollectorService
}

func NewImageService(ir ImageRepository, cs CollectorService) *imageService {
	return &imageService{
		imageRepository:  ir,
		collectorService: cs,
	}
}

func (s imageService) GetAll(ctx context.Context) ([]entity.Image, error) {
	return s.imageRepository.GetAll(ctx)
}

func (s imageService) GetByDate(ctx context.Context, date string) (entity.Image, error) {
	image, err := s.imageRepository.GetByDate(ctx, date)
	if errs.GetErrorStatus(err) != errs.NotFound {
		return image, err
	}

	params := dto.ApodQueryParams{
		Date: date,
	}

	image, err = s.collectorService.GetAPOD(ctx, params)
	if err != nil {
		return entity.Image{}, fmt.Errorf("get APOD: %w", err)
	}

	return s.imageRepository.Create(ctx, image)
}
