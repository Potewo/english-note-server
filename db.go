package main

import (
	"math"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Note struct {
	gorm.Model
	English     string
	Japanese    string
	Description string
	Examples    string
	Similar     string
	Tags        []Tag
}

type Record struct {
	gorm.Model
	Correct uint
	NoteID  uint
}

type Tag struct {
	gorm.Model
	NoteID uint
	Name   string
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

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (d *DB) AddNote(notes []Note) ([]Note, error) {
	c := d.db.Create(&notes)
	if c.Error != nil {
		return nil, c.Error
	}
	return notes, nil
}

func (d *DB) ReadAllNotes(page int, pageSize int) ([]Note, int, error) {
	notes := []Note{}
	totalItems := d.db.Find(&notes).RowsAffected
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	c := d.db.Scopes(Paginate(page, pageSize)).Preload(clause.Associations).Find(&notes)
	if c.Error != nil {
		return nil, 0, c.Error
	}
	return notes, totalPages, nil
}

func (d *DB) ReadNote(id uint) (Note, error) {
	note := Note{}
	c := d.db.Preload(clause.Associations).First(&note, id)
	if c.Error != nil {
		return Note{}, c.Error
	}
	return note, nil
}

func (d *DB) UpdateNotes(notes []Note) ([]Note, error) {
	for i := range notes {
		note, err := d.ReadNote(notes[i].ID)
		if err != nil {
			return nil, err
		}
		for _, savedTag := range note.Tags {
			have := false
			for _, requestedTag := range notes[i].Tags {
				if savedTag.ID == requestedTag.ID {
					have = true
					break
				}
			}
			if !have {
				err = d.DeleteTags([]Tag{savedTag})
				if err != nil {
					return nil, err
				}
			}
		}
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
