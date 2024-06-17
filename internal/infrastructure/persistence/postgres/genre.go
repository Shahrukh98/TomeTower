package postgres

import (
	"database/sql"
	"tometower/internal/domain/genre"
)

type GenrePostgresRepository struct {
	db *sql.DB
}

func NewGenrePostgresRepository(db *sql.DB) *GenrePostgresRepository {
	return &GenrePostgresRepository{db: db}
}

func (r *GenrePostgresRepository) GetAllGenres() ([]genre.Genre, error) {
	rows, err := r.db.Query("select * from genres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []genre.Genre
	for rows.Next() {
		var genre genre.Genre
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil

}

func (r *GenrePostgresRepository) GetByID(id string) (genre.Genre, error) {
	var genre genre.Genre

	row := r.db.QueryRow("SELECT * FROM genres WHERE id = $1", id)
	err := row.Scan(&genre.ID, &genre.Name, &genre.CreatedAt, &genre.UpdatedAt)
	return genre, err
}
