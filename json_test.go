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
	err = ReadJson("notExists.json", &s)
	if err != nil {
		t.Errorf("Error occured when read not exists json file:\n%v\n", err)
	}
	if len(s) != 0 {
		t.Errorf("Error occured when read not exits json file.returned not empty slice")
	}
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
			UUID:        "3e78def4-cce4-4cb0-bdd0-43accdac7caf",
		},
		{
			English:     "New english2",
			Japanese:    "New japanese2",
			Description: "New description2",
			Examples:    "New examples2",
			Similar:     "New similar2",
			Tags:        []string{"newtag2_1", "newtag2_2"},
			UUID:        "6d1c5cde-a77f-486e-811f-f5716d99281b",
		},
	}
	err := AppendJson("append_test.json", &newNotes)
	if err != nil {
		t.Errorf("Error occured when append json: %v", err)
	}
}
