package postgres

import (
	"database/sql"
	"log"
	"tometower/internal/entity"
)

type AuthorPostgresRepository struct {
	db *sql.DB
}

func NewAuthorPostgresRepository(db *sql.DB) *AuthorPostgresRepository {
	return &AuthorPostgresRepository{db: db}
}

func (r *AuthorPostgresRepository) GetAll() ([]entity.Author, error) {
	rows, err := r.db.Query("select * from authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entity.Author
	for rows.Next() {
		var author entity.Author
		var dateOfDeath sql.NullString
		if err := rows.Scan(&author.Id, &author.Name, &author.PhotoUrl, &author.Nationality, &author.DateOfBirth, &dateOfDeath, &author.CreatedAt, &author.UpdatedAt); err != nil {
			return nil, err
		}
		if dateOfDeath.Valid {
			author.DateOfDeath = dateOfDeath.String
		} else {
			author.DateOfDeath = ""
		}
		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil

}

func (r *AuthorPostgresRepository) GetById(id string) (entity.Author, error) {
	var author entity.Author
	var dateOfDeath sql.NullString

	row := r.db.QueryRow("SELECT * FROM authors WHERE id = $1", id)
	err := row.Scan(&author.Id, &author.Name, &author.PhotoUrl, &author.Nationality, &author.DateOfBirth, &dateOfDeath, &author.CreatedAt, &author.UpdatedAt)
	if dateOfDeath.Valid {
		author.DateOfDeath = dateOfDeath.String
	} else {
		author.DateOfDeath = ""
	}
	return author, err
}

func (r *AuthorPostgresRepository) Add(author entity.Author) (string, error) {
	lastInsertId := ""
	var err error

	if len(author.DateOfDeath) == 0 {
		log.Print("asd")
		err = r.db.QueryRow("INSERT INTO authors(name, photo_url, nationality, date_of_birth) VALUES($1, $2, $3, $4) RETURNING id", author.Name, author.PhotoUrl, author.Nationality, author.DateOfBirth).Scan(&lastInsertId)
	} else {
		log.Print("dsa")
		err = r.db.QueryRow("INSERT INTO authors(name, photo_url, nationality, date_of_birth, date_of_death) VALUES($1, $2, $3, $4, $5) RETURNING id", author.Name, author.PhotoUrl, author.Nationality, author.DateOfBirth, author.DateOfDeath).Scan(&lastInsertId)
	}

	if err != nil {
		return "", err
	}
	return lastInsertId, nil
}

func (r *AuthorPostgresRepository) Remove(id string) error {
	_, err := r.db.Exec("DELETE FROM authors where id=($1)", id)
	return err
}
