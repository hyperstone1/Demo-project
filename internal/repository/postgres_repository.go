package repository

import (
	"context"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type catPostgres struct {
	conn *pgx.Conn
}

func NewPG(conn *pgx.Conn) (*catPostgres, error) {

	return &catPostgres{conn}, nil
}

func (c *catPostgres) Create(cat model.Cat) error {
	_, err := c.conn.Exec(context.Background(), "INSERT INTO cats (id, name, age) VALUES ($1, $2, $3)", cat.Id, cat.Name, cat.Age)
	if err != nil {
		return err
	}
	return nil
}

func (c *catPostgres) Get() (model.Cats, error) {
	var cats model.Cats
	rows, err := c.conn.Query(context.Background(), "SELECT * FROM cats")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cat model.Cat
		err = rows.Scan(&cat.Id, &cat.Name, &cat.Age)
		if err != nil {
			return nil, err
		}
		cats = append(cats, &cat)
	}
	return cats, nil
}

func (c *catPostgres) Update(cat model.Cat, id string) error {
	_, err := c.conn.Exec(context.Background(), `UPDATE cats SET name = $1, age = $2 WHERE id = $3`, cat.Name, cat.Age, cat.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c *catPostgres) Delete(id string) error {
	_, err := c.conn.Exec(context.Background(), `DELETE FROM cats WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *catPostgres) GetById(id string) (*model.Cat, error) {
	cat := model.Cat{}
	err := c.conn.QueryRow(context.Background(), `SELECT * FROM cats WHERE id=$1`, id).Scan(&cat.Id, &cat.Name, &cat.Age)
	if err != nil {
		return nil, errors.Wrap(err, "rep.GetUserById error")
	}
	return &cat, nil
}
