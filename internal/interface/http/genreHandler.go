package http

import (
	"errors"
	"net/http"

	"tometower/internal/dto"
	"tometower/internal/service"
	"tometower/pkg/util"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.service.GetAllGenres()
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusOK, genres)
}

func (h *GenreHandler) AddGenre(w http.ResponseWriter, r *http.Request) {
	var createGenreDto dto.CreateGenreDto
	util.ParseJSONBody(w, r, &createGenreDto)

	id, err := h.service.AddGenre(createGenreDto)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusCreated, map[string]string{"id": id})
}

func (h *GenreHandler) GetGenreById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		util.ClientErrorResponse(w, http.StatusBadRequest, errors.New("cannot get id from url"))
		return
	}

	genre, err := h.service.GetGenreById(id)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusOK, genre)
}

func (h *GenreHandler) RemoveGenre(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		util.ClientErrorResponse(w, http.StatusBadRequest, errors.New("cannot get id from url"))
		return
	}

	err := h.service.RemoveGenre(id)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusNoContent, "")
}
