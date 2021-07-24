package model

// Folder is a model that represents folders of bookmarks that can exist in our app,
// containing bookmarks, owned by one user, and accessed by individual users and or
// users belonging to a certain organization.
type Folder struct {
	Model

	Name  string `json:"name"`
	Color string `json:"color"`

	ParentID *uint    `json:"parent_id"`
	Parent   *Folder `gorm:"association_foreignkey:ParentID" json:"parent"`
	OwnerID  uint    `json:"-"`
	Owner    *User   `gorm:"foreignkey:OwnerID" json:"owner"`
	//Organizations []Organization `gorm:"many2many:folder_organization;"`
	Bookmarks []Bookmark `gorm:"many2many:bookmark_folder;" json:"bookmarks"`
	//Users     []User     `gorm:"many2many:folder_user;" json:"users"`
}

// Add strings to the array to allow embedding that resource through the
// embed query paramter.
func FolderValidEmbeds() []string {
	return []string{"owner", "bookmarks"}
}
