package main

import (
	"os"
	"encoding/json"
)

func ReadJson(fileName string, s *[]Note) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(s)
	if err != nil {
		return err
	}
	return nil
}

func WriteJson(fileName string, s *[]Note) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(s)
	if err != nil {
		return err
	}
	return nil
}


func AppendJson(fileName string, s *[]Note) error {
	notes := []Note{}
	err := ReadJson(fileName, &notes)
	if err != nil {
		return err
	}
	notes = append(notes, *s...)
	err = WriteJson(fileName, &notes)
	if err != nil {
		return err
	}
	return nil
}
