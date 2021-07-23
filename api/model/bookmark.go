package model

// Bookmark is a model that represents the bookmarks our app can save. They are owned by one user,
// Can be in any number of folders, and can have any number of tags.
type Bookmark struct {
	Model

	Name  string  `json:"name"`
	URL   string  `json:"url"`
	Color *string `json:"color"`

	OwnerID uint     `json:"-"`
	Owner   *User    `gorm:"foreignKey:OwnerID" json:"owner"`
	Folders []Folder `gorm:"many2many:bookmark_folder;" json:"folders"`
	Tags    []Tag    `gorm:"many2many:bookmark_tag;" json:"tags"`
}

// Add strings to the array to allow embedding that resource through the
// embed query paramter.
func BookmarkValidEmbeds() []string {
	return []string{"owner", "folders"}
}
