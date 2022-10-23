package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/kingofmidas/astrology-api/internal/pkg/entity"
	errs "github.com/kingofmidas/astrology-api/internal/pkg/errors"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type imageRepository struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) *imageRepository {
	return &imageRepository{db: db}
}

func (r imageRepository) GetAll(ctx context.Context) ([]entity.Image, error) {
	images := make([]entity.Image, 0)

	err := r.db.SelectContext(ctx, &images, "SELECT * FROM images")

	return images, errs.UnexpectedError(err, "imageRepository.GetAll")
}

func (r imageRepository) GetByDate(ctx context.Context, date string) (entity.Image, error) {
	var image entity.Image

	err := r.db.GetContext(ctx, &image, "SELECT * FROM images WHERE date = $1", date)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Image{}, errs.New(errs.NotFound, "image not found")
	}

	return image, errs.UnexpectedError(err, "imageRepository.GetByDate")
}

func (r imageRepository) Create(ctx context.Context, img entity.Image) (entity.Image, error) {
	var image entity.Image

	err := r.db.QueryRowxContext(ctx, `
		INSERT INTO images (title, date, url, data) VALUES ($1, $2, $3, $4) RETURNING *`,
		img.Title, img.Date, img.URL, img.Data,
	).StructScan(&image)

	return image, errs.UnexpectedError(err, "imageRepository.Create")
}
