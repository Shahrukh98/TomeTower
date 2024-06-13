package postgres

import (
	"database/sql"
	"tometower/internal/domain/user"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Add(user user.User) error {
	_, err := r.db.Exec("INSERT INTO users(name, nick, email, password) VALUES($1, $2, $3, $4)", user.Name, user.Nick, user.Email, user.Password)
	return err
}

func (r *UserPostgresRepository) FindByID(id string) (user.User, error) {
	var user user.User

	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.NickUpdatedAt, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
