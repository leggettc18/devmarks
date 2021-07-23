package db

import (
	"strings"

	"github.com/jinzhu/gorm"

	// Blank becuase it is needed for gorm but never directly used
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// Database represents our App's Database (connection object, etc.)
type Database struct {
	*gorm.DB
}

// New sets up our Database connections and returns our App's Database object
func New(config *Config) (*Database, error) {
	db, err := gorm.Open("postgres", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}

func (db *Database) preloadEmbeds(valid []string, embeds []string) *gorm.DB {
	var instance = db.DB
	for _, embed := range embeds {
		if contains(valid, embed) {
			instance = instance.Preload(strings.Title(embed))
		}
	}
	return instance
}
