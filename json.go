package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func ReadJson(fileName string, s *[]Note) error {
	if !fileExists(fileName) {
		*s = []Note{}
		return nil
	}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error occured with opening a file in os.Open()\n")
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(s)
	if err != nil {
		fmt.Printf("Error occured with decoding a json file in json.NewDecoder().Decode()\n")
		return err
	}
	return nil
}

func WriteJson(fileName string, s *[]Note) error {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error occured with creating a file in os.Create()\n")
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(s)
	if err != nil {
		fmt.Printf("Error occured with encoding and writing json in json.NewEncoder.Encode()\n")
		return err
	}
	return nil
}


func AppendJson(fileName string, s *[]Note) error {
	notes := []Note{}
	err := ReadJson(fileName, &notes)
	if err != nil {
		fmt.Printf("Error occured with reding json in ReadJson()\n")
		return err
	}
	notes = append(notes, *s...)
	err = WriteJson(fileName, &notes)
	if err != nil {
		fmt.Printf("Error occured with writing json in WriteJson()\n")
		return err
	}
	return nil
}
