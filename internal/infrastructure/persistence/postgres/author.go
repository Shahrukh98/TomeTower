package postgres

import (
	"database/sql"
	"tometower/internal/domain/author"
)

type AuthorPostgresRepository struct {
	db *sql.DB
}

func NewAuthorPostgresRepository(db *sql.DB) *AuthorPostgresRepository {
	return &AuthorPostgresRepository{db: db}
}

func (r *AuthorPostgresRepository) GetAllAuthors() ([]author.Author, error) {
	rows, err := r.db.Query("select * from authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []author.Author
	for rows.Next() {
		var author author.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.PhotoUrl, &author.Nationality, &author.DateOfBirth, &author.DateOfDeath, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil

}

func (r *AuthorPostgresRepository) GetByID(id string) (author.Author, error) {
	var author author.Author

	row := r.db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	err := row.Scan(&author.ID, &author.Name, &author.PhotoUrl, &author.Nationality, &author.DateOfBirth, &author.DateOfDeath, &author.CreatedAt, &author.UpdatedAt)
	return author, err
}
