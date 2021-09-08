package postgres

import (
	"myprogs/cleanArc/models"
)

type User struct {
	ID       int64  `sql:"id,omitempty"`
	Username string `sql:"username"`
	Password string `sql:"password"`
}

func toPostgresUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *models.User  {
	return &models.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}