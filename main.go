package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Note struct {
	English     string   `json:"english"`
	Japanese    string   `json:"japanese"`
	Description string   `json:"description"`
	Examples    string   `json:"examples"`
	Similar     string   `json:"similar"`
	Tags        []string `json:"tags"`
	UUID        string   `json:"uuid"`
}

type Record struct {
	UUID string `json:"uuid"`
	Correct bool `json:"correct"`
	Date string `json:"date"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/upload", handleAdd)
	e.PUT("/upload", handleUpdate)
	e.GET("/get", handleGet)
	e.DELETE("/upload", handleDelete)

	e.Logger.Fatal(e.Start(":1323"))
}

func handleAdd(c echo.Context) error {
	u := Note{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	ul := []Note{u}
	fmt.Printf("Json received:\n%#v\n", u)

	if err := AppendJson("saveFile.json", &ul); err != nil {
		handleError(c, err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func handleGet(c echo.Context) error {
	notes := []Note{}
	err := ReadJson("saveFile.json", &notes)
	if err != nil {
		handleError(c, err)
		return err
	}
	return c.JSON(http.StatusCreated, &notes)
}

func handleUpdate(c echo.Context) error {
	changedNote := Note{}
	if err := c.Bind(&changedNote); err != nil {
		handleError(c, err)
		return err
	}
	fmt.Printf("updating: %#v\n", changedNote)
	notes := []Note{}
	err := ReadJson("saveFile.json", &notes)
	if err != nil {
		handleError(c, err)
		return err
	}
	target_i := -1
	for i, note := range notes {
		if note.UUID == changedNote.UUID {
			target_i = i
			break
		}
	}
	if target_i == -1 {
		err = errors.New("Target UUID is not found")
		handleError(c, err)
		return err
	}
	notes[target_i] = changedNote
	if err = WriteJson("saveFile.json", &notes); err != nil {
		handleError(c, err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func handleDelete(c echo.Context) error {
	deleteNote := Note{}
	if err := c.Bind(&deleteNote); err != nil {
		handleError(c, err)
		return err
	}
	fmt.Printf("deleting: %#v\n", deleteNote)
	notes := []Note{}
	err := ReadJson("saveFile.json", &notes)
	if err != nil {
		handleError(c, err)
		return err
	}
	target_i := -1
	for i, note := range notes {
		if note.UUID == deleteNote.UUID {
			target_i = i
			break
		}
	}
	if target_i == -1 {
		err = errors.New("Target UUID is not found")
		handleError(c, err)
		return err
	}
	//remove
	notes = notes[:target_i+copy(notes[target_i:], notes[target_i+1:])]
	if err = WriteJson("saveFile.json", &notes); err != nil {
		handleError(c, err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

// func handleGetRecord(c echo.Context) error {
// 	newRecords := []Record{}
// 	if err := c.Bind(&newRecords); err != nil {
// 		handleError(c, err)
// 		return err
// 	}
// 	fmt.Printf("Get record: %#v\n", newRecords)
// 	if err := AppendJson("saveRecordFile.json", &newRecords); err != nil {
// 		handleError(c, err)
// 		return err
// 	}
// 	return c.NoContent(http.StatusOK)
// }

func handleError(c echo.Context, e error) {
	fmt.Printf("Error:\n%#v\n", e)
	sendingErr := c.JSON(http.StatusInternalServerError, e)
	if sendingErr != nil {
		fmt.Printf("Error occured with sending json error message")
	}
}
