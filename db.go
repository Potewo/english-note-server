package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Note struct {
	gorm.Model
	English string
	Japanese string
	Description string
	Examples string
	Similar string
	Tags []Tag
}

type Record struct {
	gorm.Model
	Correct bool
	NoteID uint
}

type Tag struct {
	gorm.Model
	NoteID uint
	Name string
}

type DB struct {
	db *gorm.DB
}

func (d *DB) gormConnect(name string) error {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func NewDB(name string) (*DB, error) {
	d := new(DB)
	err := d.gormConnect(name)
	if err != nil {
		return nil, err
	}
	err = d.db.AutoMigrate(&Note{}, &Record{}, &Tag{})
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DB) AddNote(notes []Note) ([]Note, error) {
	c := d.db.Create(&notes)
	if c.Error != nil {
		return nil, c.Error
	}
	return notes, nil
}

func (d *DB) ReadAllNotes() ([]Note, error) {
	notes := []Note{}
	c := d.db.Preload(clause.Associations).Find(&notes)
	if c.Error != nil {
		return nil, c.Error
	}
	return notes, nil
}

func (d *DB) UpdateNotes(notes []Note) ([]Note, error) {
	for i := range notes {
		c := d.db.Model(&notes[i]).Select("*").Updates(&notes[i])
		if c.Error != nil {
			return nil, c.Error
		}
	}
	return notes, nil
}

func (d *DB) DeleteNotes(notes []Note) error {
	for _, note := range notes {
		c := d.db.Delete(&note)
		if c.Error != nil {
			return c.Error
		}
	}
	return nil
}

func (d *DB) UpdateTags(tags []Tag) error {
	for _, tag := range tags {
		c := d.db.Model(&tag).Select("*").Updates(tag)
		if c.Error != nil {
			return c.Error
		}
	}
	return nil
}

func (d *DB) DeleteTags(tags []Tag) error {
	for _, tag := range tags {
		c := d.db.Delete(&tag)
		if c.Error != nil {
			return c.Error
		}
	}
	return nil
}

func (d *DB) ReadAllRecords() ([]Record, error) {
	records := []Record{}
	c := d.db.Find(&records)
	if c.Error != nil {
		return nil, c.Error
	}
	return records, nil
}

func (d *DB) AddRecords(records []Record) ([]Record, error) {
	c := d.db.Create(&records)
	if c.Error != nil {
		return nil, c.Error
	}
	return records, nil
}
