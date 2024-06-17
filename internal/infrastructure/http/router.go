package http

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"tometower/internal/domain/genre"
	"tometower/internal/domain/user"
	"tometower/internal/infrastructure/persistence/postgres"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type UserClaims struct {
	ID   string
	Name string
	jwt.RegisteredClaims
}

func GetUserFromContext(r *http.Request) *UserClaims {
	if user, ok := r.Context().Value("props").(*UserClaims); ok {
		return user
	}
	return nil
}

func VerifyToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		user, err := VerifyToken(token)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Store user info in context
		ctx := context.WithValue(r.Context(), "props", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenreRouter(repo genre.GenreRepository) *http.ServeMux {
	service := genre.NewGenreService(repo)
	handler := NewGenreHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("/", handler.GetAll)
	router.HandleFunc("/{genre}", handler.GetByID)
	return router
}

func UserRouter(repo user.UserRepository) *http.ServeMux {
	service := user.NewUserService(repo)
	handler := NewUserHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("/register", handler.AddUser)
	router.HandleFunc("/login", handler.GetByEmail)
	router.Handle("/update-nick", AuthMiddleware(http.HandlerFunc(handler.UpdateNick)))
	return router
}

func TomeTowerRouter(db *sql.DB) *http.ServeMux {
	userRepo := postgres.NewUserPostgresRepository(db)
	genreRepo := postgres.NewGenrePostgresRepository(db)

	userRouter := UserRouter(userRepo)
	genreRotuer := GenreRouter(genreRepo)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/users/", http.StripPrefix("/users", userRouter))
	mainRouter.Handle("/genres/", http.StripPrefix("/genres", genreRotuer))

	return mainRouter
}
