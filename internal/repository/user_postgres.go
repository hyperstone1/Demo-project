package repository

import (
	"context"

	"github.com/hyperstone1/TestCRUD/internal/model"
)

func (c *catPostgres) CreateUser(user model.User) (int, error) {
	var id int
	err := c.conn.QueryRow(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *catPostgres) GetUser(username, password string) (model.User, error) {
	var err error
	var user model.User

	query := c.conn.QueryRow(context.Background(), "SELECT username,password FROM users WHERE username=$1 AND password=$2", username, password)

	err = query.Scan(&user.Username, &user.Password)
	return user, err

}
