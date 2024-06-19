package author

type AuthorService struct {
	repo AuthorRepository
}

func NewAuthorService(repo AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]Author, error) {
	return s.repo.GetAllAuthors()
}

func (s *AuthorService) GetByID(id string) (Author, error) {
	return s.repo.GetByID(id)
}
