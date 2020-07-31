package db

import (
	"github.com/pkg/errors"

	"leggett.dev/devmarks/api/model"
)

func (db *Database) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	return &user, errors.Wrap(db.First(&user, model.User{Email: email}).Error, "unable to get user")
}

func (db *Database) CreateUser(user *model.User) error {
	return db.Create(user).Error
}
