package db

import (
	"github.com/pkg/errors"
	"leggett.dev/devmarks/api/model"
)

func (db *Database) CreateFolder(folder *model.Folder) error {
	return errors.Wrap(db.Create(folder).Error, "unable to create folder")
}

func (db *Database) GetFolderByID(id uint) (*model.Folder, error) {
	var folder model.Folder
	return &folder, errors.Wrap(db.Preload("Bookmarks").First(&folder, id).Error, "unable to get folder")
}

func (db *Database) AddBookmarkToFolder(bookmark_id uint, folder_id uint) error {
	bookmark, err := db.GetBookmarkByID(bookmark_id, nil)
	if err != nil {
		return err
	}
	folder, err := db.GetFolderByID(folder_id)
	if err != nil {
		return err
	}
	err = errors.Wrap(db.Model(&folder).Association("Bookmarks").Append([]model.Bookmark{*bookmark}).Error, "unable to add bookmark to folder")
	if err != nil {
		return err
	}
	return nil
}
