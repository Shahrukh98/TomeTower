package postgres

import (
	"database/sql"
	"tometower/internal/entity"
)

type GenrePostgresRepository struct {
	db *sql.DB
}

func NewGenrePostgresRepository(db *sql.DB) *GenrePostgresRepository {
	return &GenrePostgresRepository{db: db}
}

func (r *GenrePostgresRepository) GetAll() ([]entity.Genre, error) {
	rows, err := r.db.Query("select * from genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []entity.Genre
	for rows.Next() {
		var genre entity.Genre
		if err := rows.Scan(&genre.Id, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil

}

func (r *GenrePostgresRepository) GetById(id string) (entity.Genre, error) {
	var genre entity.Genre

	row := r.db.QueryRow("SELECT * FROM genres WHERE id = $1", id)
	err := row.Scan(&genre.Id, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
	return genre, err
}

func (r *GenrePostgresRepository) Add(genre entity.Genre) (string, error) {
	lastInsertId := ""
	err := r.db.QueryRow("INSERT INTO genres(name) VALUES($1) RETURNING id", genre.Name).Scan(&lastInsertId)
	if err != nil {
		return "", err
	}
	return lastInsertId, nil
}

func (r *GenrePostgresRepository) Remove(id string) error {
	_, err := r.db.Exec("DELETE FROM genres where id=($1)", id)
	return err
}
