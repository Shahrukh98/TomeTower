package http

import (
	"database/sql"
	"net/http"

	"tometower/internal/domain/user"
	"tometower/internal/infrastructure/persistence/postgres"
)

func UserRouter(repo user.UserRepository) *http.ServeMux {
	service := user.NewUserService(repo)
	handler := NewUserHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("/register", handler.AddUser)
	router.HandleFunc("/login", handler.FindByEmail)
	router.HandleFunc("/update-nick", handler.UpdateNick)
	return router
}

func TomeTowerRouter(db *sql.DB) *http.ServeMux {
	repo := postgres.NewUserPostgresRepository(db)

	userRouter := UserRouter(repo)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/users/", http.StripPrefix("/users", userRouter))

	return mainRouter
}
