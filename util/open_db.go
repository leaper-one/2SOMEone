package util

import (
	"log"

	"github.com/leaper-one/2SOMEone/core"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	write *gorm.DB
	read  *gorm.DB
}

func OpenDB(path string) *DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	db.AutoMigrate(&core.BasicUser{}, &core.Note{}, &core.Message{}, &core.BiliUser{}, &core.Message{})
	return &DB{write: db, read: db}
}

func (db *DB) Update() *gorm.DB {
	return db.write
}

func (db *DB) View() *gorm.DB {
	return db.read
}

func (db *DB) Debug() *DB {
	return &DB{
		write: db.write.Debug(),
		read:  db.read.Debug(),
	}
}

func (db *DB) Begin() *DB {
	tx := db.write.Begin()
	return &DB{
		write: tx,
		read:  db.read,
	}
}

func (db *DB) Rollback() error {
	return db.write.Rollback().Error
}

func (db *DB) Commit() error {
	return db.write.Commit().Error
}

func (db *DB) RollbackUnlessCommitted() {
	if err := db.write.Rollback().Error; err != nil {
		logrus.WithError(err).Error("DB: RollbackUnlessCommitted")
	}
}

func (db *DB) Tx(fn func(tx *DB) error) error {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
