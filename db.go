package ovpm

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// We blank import sqlite here because gorm needs it.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *DB

// DB represents a persistent storage.
type DB struct {
	*gorm.DB
}

// CreateDB prepares and returns new storage.
//
// It should be run at the start of the program.
func CreateDB(dialect string, args ...interface{}) *DB {
	if len(args) > 0 && args[0] == "" {
		args[0] = _DefaultDBPath
	}
	var err error

	dbase, err := gorm.Open(dialect, args...)
	if err != nil {
		logrus.Fatalf("couldn't open sqlite database %v: %v", args, err)
	}

	dbase.AutoMigrate(&dbUserModel{})
	dbase.AutoMigrate(&dbServerModel{})
	dbase.AutoMigrate(&dbRevokedModel{})
	dbase.AutoMigrate(&dbNetworkModel{})

	dbPTR := &DB{DB: dbase}
	db = dbPTR
	return dbPTR
}

// Cease closes the database.
//
// It should be run at the exit of the program.
func (db *DB) Cease() {
	db.DB.Close()
}

// GetUserIPOutByUsernames returns a username->ip_out map for given users.
//
// It prefers the already initialized global db instance; if not available,
// it opens a temporary sqlite connection with the default db path.
func GetUserIPOutByUsernames(usernames []string) (map[string]string, error) {
	out := map[string]string{}
	if len(usernames) == 0 {
		return out, nil
	}

	type userIPOutRow struct {
		Username string
		IPOut    string `gorm:"column:ip_out"`
	}

	queryDB := db
	var tmpDB *gorm.DB
	if queryDB == nil {
		var err error
		tmpDB, err = gorm.Open("sqlite3", _DefaultDBPath)
		if err != nil {
			return nil, err
		}
		defer tmpDB.Close()
		queryDB = &DB{DB: tmpDB}
	}

	var rows []userIPOutRow
	q := queryDB.Table("db_user_models").
		Select("username, ip_out").
		Where("username IN (?)", usernames).
		Scan(&rows)
	if q.Error != nil {
		return nil, q.Error
	}

	for _, r := range rows {
		out[r.Username] = r.IPOut
	}
	return out, nil
}
