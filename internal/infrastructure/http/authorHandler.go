package http

import (
	"net/http"
	"strings"

	"tometower/internal/domain/author"
	"tometower/pkg/util"
)

type AuthorHandler struct {
	service *author.AuthorService
}

func NewAuthorHandler(service *author.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			authors, err := h.service.GetAllAuthors()
			if err != nil {
				util.LogError(err, "Failed to get authors")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get authors"})
				return
			}

			util.JSONResponse(w, http.StatusOK, authors)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			url_splits := strings.Split(r.URL.Path, "/")
			if len(url_splits) != 2 {
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get author id"})
			}

			id := string(url_splits[1])
			author, err := h.service.GetByID(id)
			if err != nil {
				util.LogError(err, "Failed to get authors")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get authors"})
				return
			}

			util.JSONResponse(w, http.StatusOK, author)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
