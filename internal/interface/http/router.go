package http

import (
	"database/sql"
	"net/http"

	"tometower/internal/middleware"
	"tometower/internal/persistence/postgres"
	"tometower/internal/repository"
	"tometower/internal/service"
)

func AuthRouter(userService service.UserService) *http.ServeMux {
	authService := service.NewAuthService(userService)
	handler := NewAuthHandler(authService)

	router := http.NewServeMux()
	router.HandleFunc("POST /register", handler.Register)
	router.HandleFunc("POST /login", handler.Login)
	return router
}

func AuthorRouter(repo repository.AuthorRepository) *http.ServeMux {
	service := service.NewAuthorService(repo)
	handler := NewAuthorHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handler.GetAllAuthors)
	router.HandleFunc("POST /", handler.AddAuthor)
	router.HandleFunc("GET /{id}", handler.GetAuthorById)
	router.HandleFunc("DELETE /{id}", handler.RemoveAuthor)
	return router
}

func GenreRouter(repo repository.GenreRepository) *http.ServeMux {
	service := service.NewGenreService(repo)
	handler := NewGenreHandler(service)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handler.GetAllGenres)
	router.HandleFunc("GET /{id}", handler.GetGenreById)
	router.HandleFunc("POST /", handler.AddGenre)
	router.HandleFunc("DELETE /{id}", handler.RemoveGenre)
	return router
}

func UserRouter(userService service.UserService) *http.ServeMux {
	handler := NewUserHandler(&userService)

	router := http.NewServeMux()
	router.HandleFunc("GET /{id}", http.HandlerFunc(handler.GetUserById))
	router.Handle("PATCH /update-nick", middleware.AuthMiddleware(http.HandlerFunc(handler.UpdateNick)))
	return router
}

func Router(db *sql.DB) *middleware.Logger {
	userRepo := postgres.NewUserPostgresRepository(db)
	genreRepo := postgres.NewGenrePostgresRepository(db)
	authorRepo := postgres.NewAuthorPostgresRepository(db)

	userService := service.NewUserService(userRepo)

	authRouter := AuthRouter(*userService)
	userRouter := UserRouter(*userService)
	genreRouter := GenreRouter(genreRepo)
	authorRouter := AuthorRouter(authorRepo)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", authRouter))
	mainRouter.Handle("/users/", http.StripPrefix("/users", userRouter))
	mainRouter.Handle("/genres/", http.StripPrefix("/genres", genreRouter))
	mainRouter.Handle("/authors/", http.StripPrefix("/authors", authorRouter))

	routerWithLogger := middleware.NewLogger(mainRouter)
	return routerWithLogger
}
