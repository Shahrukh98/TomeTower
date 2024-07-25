package http

import (
	"errors"
	"net/http"

	"tometower/internal/dto"
	"tometower/internal/middleware"
	"tometower/internal/service"
	"tometower/pkg/util"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		util.ClientErrorResponse(w, http.StatusBadRequest, errors.New("cannot get id from url"))
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateNick(w http.ResponseWriter, r *http.Request) {
	var parsedJson dto.NickUpdateDto
	err := util.ParseJSONBody(w, r, &parsedJson)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	userC := middleware.GetUserFromContext(r)

	err = h.service.UpdateNick(userC.ID, parsedJson.Nick)
	if err != nil {
		util.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	util.JSONResponse(w, http.StatusNoContent, "")
}
