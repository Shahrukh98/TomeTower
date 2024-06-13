package http

import (
	"net/http"
	"time"
	"tometower/internal/domain/user"
	"tometower/pkg/util"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var user user.User
			err := util.ParseJSONBody(w, r, &user)
			if err != nil {
				util.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
				return
			}

			err = h.service.AddUser(user)
			if err != nil {
				util.LogError(err, "Failed to add user")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add user"})
				return
			}

			util.JSONResponse(w, http.StatusCreated, map[string]string{"status": "User added successfully"})
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			id := r.URL.Query().Get("id")
			if id == "" {
				util.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing id parameter"})
				return
			}

			user, err := h.service.FindUserById(id)
			if err != nil {
				util.LogError(err, "Failed to get user")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
				return
			}

			util.JSONResponse(w, http.StatusOK, user)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) UpdateNick(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			id := r.URL.Query().Get("id")
			if id == "" {
				util.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing id parameter"})
				return
			}

			userObj, err := h.service.FindUserById(id)
			if err != nil {
				util.LogError(err, "Failed to get user")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
				return
			}

			currentTime := time.Now().Unix()

			timeDiff := currentTime - userObj.NickUpdatedAt.Unix()
			if timeDiff < user.NickUpdateCooldown {
				util.JSONResponse(w, http.StatusConflict, map[string]string{"status": "User Nick Update on Cooldown!"})
			} else {
				var parsedJson user.NickUpdate
				err = util.ParseJSONBody(w, r, &parsedJson)
				err = h.service.UpdateNick(id, parsedJson.Nick)
				if err != nil {
					util.LogError(err, "Failed to get useraa")
					util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
					return
				}

				util.JSONResponse(w, http.StatusNoContent, map[string]string{"status": "User Nick Updated successfully"})
			}
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
