package main

import (
	"fmt"
	"testing"
)

func TestReadJson(t *testing.T) {
	s := []Note{}
	err := ReadJson("sample.json", &s)
	if err != nil {
		t.Errorf("Error occured when read json: %v", err)
	}
	fmt.Printf("%#v\n", s)
}

func TestWriteJson(t *testing.T) {
	s := []Note{}
	err := ReadJson("sample.json", &s)
	if err != nil {
		t.Errorf("Error occured when read json: %v", err)
	}
	s = append(s, Note{
		English:     "New english",
		Japanese:    "New japanese",
		Description: "New description",
		Examples:    "New examples",
		Similar:     "New similar",
		Tags:        []string{"newtag1", "newtag2"},
	})

	err = WriteJson("write_test.json", &s)
	if err != nil {
		t.Errorf("Error occured when read json: %v", err)
	}
}

func TestAppendJson(t *testing.T) {
	newNotes := []Note{
		{
			English:     "New english",
			Japanese:    "New japanese",
			Description: "New description",
			Examples:    "New examples",
			Similar:     "New similar",
			Tags:        []string{"newtag1", "newtag2"},
		},
		{
			English:     "New english2",
			Japanese:    "New japanese2",
			Description: "New description2",
			Examples:    "New examples2",
			Similar:     "New similar2",
			Tags:        []string{"newtag2_1", "newtag2_2"},
		},
	}
	err := AppendJson("append_test.json", &newNotes)
	if err != nil {
		t.Errorf("Error occured when append json: %v", err)
	}
}
