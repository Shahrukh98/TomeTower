package http

import (
	"net/http"

	"tometower/internal/dto"
	"tometower/internal/service"
	"tometower/pkg/util"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var createUserDto dto.UserCreateDto
			err := util.ParseJSONBody(w, r, &createUserDto)
			if err != nil {
				util.ClientErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			err = h.service.Register(createUserDto)
			if err != nil {
				util.ClientErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			util.JSONResponse(w, http.StatusCreated, map[string]string{"status": "User added successfully"})
		}
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var userLogin dto.UserLoginDto
			err := util.ParseJSONBody(w, r, &userLogin)
			if err != nil {
				util.ClientErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			user, err := h.service.Login(userLogin.Username, userLogin.Password)
			if err != nil {
				util.ClientErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			util.JSONResponse(w, http.StatusOK, user)
		}
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
