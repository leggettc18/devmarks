package db

import (
	"context"
	
	"github.com/pkg/errors"
	"leggett.dev/devmarks/api/helpers"
	"leggett.dev/devmarks/api/model"
)

func (db *Database) CreateFolder(folder *model.Folder) error {
	return errors.Wrap(db.Create(folder).Error, "unable to create folder")
}

func (db *Database) GetFoldersByUserID(ctx context.Context, userID uint) ([]*model.Folder, error) {
	var folders []*model.Folder
	embeds, ok := ctx.Value(helpers.EmbedsKey).([]string)
	if !ok {
		return nil, errors.New("embeds parsing error")
	}
	return folders, errors.Wrap(db.preloadEmbeds(model.FolderValidEmbeds(), embeds).Find(&folders, model.Folder{OwnerID: userID}).Error, "unable to get folders")
}

func (db *Database) GetFolderByID(ctx context.Context, id uint) (*model.Folder, error) {
	var folder model.Folder
	embeds, ok := ctx.Value(helpers.EmbedsKey).([]string)
	if !ok {
		return nil, errors.New("embeds parsing error")
	}
	return &folder, errors.Wrap(db.preloadEmbeds(model.FolderValidEmbeds(), embeds).First(&folder, id).Error, "unable to get folder")
}

func (db *Database) AddBookmarkToFolder(ctx context.Context, bookmark_id uint, folder_id uint) error {
	bookmark, err := db.GetBookmarkByID(ctx, bookmark_id)
	if err != nil {
		return err
	}
	folder, err := db.GetFolderByID(ctx, folder_id)
	if err != nil {
		return err
	}
	err = errors.Wrap(db.Model(&folder).Association("Bookmarks").Append([]model.Bookmark{*bookmark}).Error, "unable to add bookmark to folder")
	if err != nil {
		return err
	}
	return nil
}
