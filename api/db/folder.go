package db

import (
	"github.com/pkg/errors"
	"leggett.dev/devmarks/api/model"
)

func (db *Database) CreateFolder(folder *model.Folder) error {
	return errors.Wrap(db.Create(folder).Error, "unable to create folder")
}
