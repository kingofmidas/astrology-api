package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kingofmidas/astrology-api/internal/app/api/helpers"
	"github.com/kingofmidas/astrology-api/internal/pkg/dto"
	errs "github.com/kingofmidas/astrology-api/internal/pkg/errors"
)

type imageHandler struct {
	imageService ImageService
}

func NewImageHandler(is ImageService) *imageHandler {
	return &imageHandler{
		imageService: is,
	}
}

func (h *imageHandler) InitRoutes(router *mux.Router) {
	r := router.PathPrefix("/images").Subrouter()

	r.HandleFunc("", h.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{date}", h.GetByDate).Methods(http.MethodGet)
}

func (h imageHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	images, err := h.imageService.GetAll(r.Context())
	if err != nil {
		helpers.Fail(w, err)
		return
	}

	helpers.Success(w, http.StatusOK, dto.GetImagesResponse{
		Images: images,
	})
}

func (h imageHandler) GetByDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		helpers.Fail(w, errs.New(errs.BadRequest, "invalid date format"))
		return
	}

	image, err := h.imageService.GetByDate(r.Context(), date)
	if err != nil {
		helpers.Fail(w, err)
		return
	}

	helpers.Success(w, http.StatusOK, image)
}
