package dto

import "github.com/kingofmidas/astrology-api/internal/pkg/entity"

type GetImagesResponse struct {
	Images []entity.Image `json:"images"`
}
