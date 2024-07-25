package http

import (
	"errors"
	"net/http"

	"tometower/internal/dto"
	"tometower/internal/service"
	"tometower/pkg/util"
)

type AuthorHandler struct {
	service *service.AuthorService
}

func NewAuthorHandler(service *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusOK, authors)
}

func (h *AuthorHandler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var createAuthorDto dto.CreateAuthorDto
	util.ParseJSONBody(w, r, &createAuthorDto)

	id, err := h.service.AddAuthor(createAuthorDto)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	util.JSONResponse(w, http.StatusCreated, map[string]string{"id": id})
}

func (h *AuthorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		util.ClientErrorResponse(w, http.StatusBadRequest, errors.New("cannot get id from url"))
		return
	}

	author, err := h.service.GetById(id)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
	}
	util.JSONResponse(w, http.StatusOK, author)
}

func (h *AuthorHandler) RemoveAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		util.ClientErrorResponse(w, http.StatusBadRequest, errors.New("cannot get id from url"))
		return
	}

	err := h.service.RemoveAuthor(id)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
	}
	util.JSONResponse(w, http.StatusNoContent, "")
}
