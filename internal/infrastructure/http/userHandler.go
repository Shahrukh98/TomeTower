package http

import (
	"log"
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

func (h *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var userRequest user.UserLoginRequest
			var userResponse user.UserLoginResponse

			err := util.ParseJSONBody(w, r, &userRequest)
			if err != nil {
				util.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
				return
			}

			user, err := h.service.FindByEmail(userRequest.Email)
			if err != nil {
				util.LogError(err, "Failed to get user by email")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get user by email"})
				return
			}

			err = util.VerifyPassword(user.Password, userRequest.Password)
			if err != nil {
				util.LogError(err, "Invalid Password")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Invalid Password"})
				return
			}

			token, err := util.CreateToken(user.ID, user.Name)

			if err != nil {
				util.LogError(err, "Cant create token")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Cant create token"})
				return
			}

			userResponse.ID = user.ID
			userResponse.Name = user.Name
			userResponse.Nick = user.Nick
			userResponse.Token = token

			util.JSONResponse(w, http.StatusOK, userResponse)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			var userResponse user.UserLoginResponse
			id := r.URL.Query().Get("id")
			if id == "" {
				util.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Missing id parameter"})
				return
			}

			user, err := h.service.FindUserById(id)
			token, err := util.CreateToken(user.ID, user.Name)

			if err != nil {
				util.LogError(err, "Failed to get user")
				util.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
				return
			}
			userResponse.ID = user.ID
			userResponse.Name = user.Name
			userResponse.Token = token

			util.JSONResponse(w, http.StatusOK, userResponse)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) UpdateNick(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			userC := GetUserFromContext(r)
			log.Println(userC.ID) // Just testing
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
