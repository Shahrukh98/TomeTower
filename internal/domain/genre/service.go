package genre

type GenreService struct {
	repo GenreRepository
}

func NewGenreService(repo GenreRepository) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) GetAllGenres() ([]Genre, error) {
	return s.repo.GetAllGenres()
}

func (s *GenreService) GetByID(id string) (Genre, error) {
	return s.repo.GetByID(id)
}
