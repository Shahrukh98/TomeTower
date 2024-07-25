package postgres

import (
	"database/sql"
	"tometower/internal/entity"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Add(user entity.User) error {
	_, err := r.db.Exec("INSERT INTO users(username, nick, email, password, role) VALUES($1, $2, $3, $4, $5)", user.Name, user.Nick, user.Email, user.Password, user.Role)
	return err
}

func (r *UserPostgresRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User

	row := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email)
	err := row.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.Role, &user.NickUpdatedAt, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *UserPostgresRepository) GetById(id string) (entity.User, error) {
	var user entity.User

	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.Role, &user.NickUpdatedAt, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *UserPostgresRepository) UpdateNick(id string, nick string) error {
	_, err := r.db.Exec("update users set nick = $1, nick_last_updated = current_timestamp, updated_at = current_timestamp  where id = $2", nick, id)
	return err
}
