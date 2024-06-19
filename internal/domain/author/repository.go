package author

type AuthorRepository interface {
	GetAllAuthors() ([]Author, error)
	GetByID(id string) (Author, error)
}
