package db

import (
	"context"

	"github.com/pkg/errors"
	"leggett.dev/devmarks/api/model"
)

func (db *Database) CreateFolder(folder *model.Folder) error {
	return errors.Wrap(db.Create(folder).Error, "unable to create folder")
}

func (db *Database) GetFoldersByUserID(userID uint) ([]*model.Folder, error) {
	var folders []*model.Folder
	return folders, errors.Wrap(db.Find(&folders, model.Folder{OwnerID: userID}).Error, "unable to get folders")
}

func (db *Database) GetFolderByID(id uint) (*model.Folder, error) {
	var folder model.Folder
	return &folder, errors.Wrap(db.Preload("Bookmarks").First(&folder, id).Error, "unable to get folder")
}

func (db *Database) AddBookmarkToFolder(ctx context.Context, bookmark_id uint, folder_id uint) error {
	bookmark, err := db.GetBookmarkByID(ctx, bookmark_id)
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
