package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kingofmidas/astrology-api/internal/pkg/dto"
	"github.com/kingofmidas/astrology-api/internal/pkg/entity"
	errs "github.com/kingofmidas/astrology-api/internal/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	apiURL = "https://api.nasa.gov/planetary/apod"
)

type collectorService struct {
	nasaAPIKey      string
	imageRepository ImageRepository
}

func NewCollectorService(apiKey string, ir ImageRepository) collectorService {
	return collectorService{
		nasaAPIKey:      apiKey,
		imageRepository: ir,
	}
}

func (s collectorService) Collect(ctx context.Context) {
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			logrus.Info("getting APOD...")

			image, err := s.GetAPOD(ctx, dto.ApodQueryParams{})
			if err != nil {
				logrus.Errorf("get APOD: %v", err)
			}

			_, err = s.imageRepository.GetByDate(ctx, image.Date)

			if errs.GetErrorStatus(err) == errs.NotFound {
				_, err = s.imageRepository.Create(ctx, image)
				if err != nil {
					logrus.Errorf("save image: %v", err)
				}
			}
		}
	}
}

func (s collectorService) GetAPOD(ctx context.Context, params dto.ApodQueryParams) (entity.Image, error) {
	url := fmt.Sprintf("%s?api_key=%s", apiURL, s.nasaAPIKey)

	if params.Date != "" {
		url = fmt.Sprintf("%s&date=%s", url, params.Date)
	}

	resp, err := http.Get(url)
	if err != nil {
		return entity.Image{}, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return entity.Image{}, fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	var data dto.ApodResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return entity.Image{}, fmt.Errorf("unmarshall response body: %w", err)
	}

	if data.MediaType == "Video" {
		return entity.Image{}, errors.New("incorrect media type")
	}

	img, err := downloadImage(data.URL)
	if err != nil {
		return entity.Image{}, fmt.Errorf("download image: %w", err)
	}

	image := entity.Image{
		Title: data.Title,
		Date:  data.Date,
		URL:   data.URL,
		Data:  img,
	}

	return image, nil
}

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	var buff bytes.Buffer
	_, err = io.Copy(&buff, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("copy body to buffer: %w", err)
	}

	return buff.Bytes(), nil
}
