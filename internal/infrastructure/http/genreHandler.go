package http

import (
	"net/http"
	"strings"

	"tometower/internal/domain/genre"
	"tometower/pkg/util"
)

type GenreHandler struct {
	service *genre.GenreService
}

func NewGenreHandler(service *genre.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			genres, err := h.service.GetAllGenres()
			if err != nil {
				util.LogError(err, "Failed to get genres")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get genres"})
				return
			}

			util.JSONResponse(w, http.StatusOK, genres)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GenreHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			url_splits := strings.Split(r.URL.Path, "/")
			if len(url_splits) != 2 {
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get genre id"})
			}

			id := string(url_splits[1])
			genre, err := h.service.GetByID(id)
			if err != nil {
				util.LogError(err, "Failed to get genres")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get genres"})
				return
			}

			util.JSONResponse(w, http.StatusOK, genre)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
