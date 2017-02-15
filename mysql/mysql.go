package mysql

import (
	"database/sql"

	"github.com/nametake/go-example-myapp"
)

type UserService struct {
	DB                 *sql.DB
	TransactionService myapp.TransactionService
}

func (s *UserService) User(id int) (*myapp.User, error) {
	var u myapp.User
	row := s.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Users() ([]*myapp.User, error) {
	rows, err := s.DB.Query(`SELECT id, name FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*myapp.User
	for rows.Next() {
		var u myapp.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (s *UserService) CreateUser(u *myapp.User) error {
	_, err := s.DB.Exec(`INSERT INTO users ("id", "name") VALUES (?, ?)`, u.ID, u.Name)
	return err
}
