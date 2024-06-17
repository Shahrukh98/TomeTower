package genre

type GenreRepository interface {
	GetAllGenres() ([]Genre, error)
	GetByID(id string) (Genre, error)
}
