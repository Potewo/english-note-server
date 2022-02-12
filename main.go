package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *DB

func main() {
	var err error
	db, err = NewDB("saveFiles/saveFile.db")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:\t${time_rfc3339}\n" +
		"remote ip:\t${remote_ip}\n" +
		"uri:\t${uri}\n" +
		"method:\t${method}\n" +
		"status:\t${status}\n" +
		"error:\t${error}\n" +
		"header:\t${header:body}\n" +
		"-------------------\n\n",
	}))
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.POST("/note", handleAddNote)
	e.PUT("/note", handleUpdateNotes)
	e.GET("/note", handleGetNotes)
	e.DELETE("/note", handleDeleteNotes)
	e.GET("/record", handleGetRecord)
	e.POST("/record", handleAddRecord)
	e.Static("/", "public")
	e.File("/", "public/index.html")
	e.File("/new/*", "public/index.html")
	e.File("/edit/*", "public/index.html")
	e.File("/notes/*", "public/index.html")

	e.Logger.Fatal(e.Start(":1323"))
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
  fmt.Printf("Request Body:\t%v\n", string(reqBody))
  fmt.Printf("Response Body:\t%v\n", string(resBody))
}

func handleAddNote(c echo.Context) error {
	newNote := []Note{}
	if err := c.Bind(&newNote); err != nil {
		return err
	}
	fmt.Printf("Json received:\n%#v\n", newNote)
	notes, err := db.AddNote(newNote)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &notes)
}

func handleGetNotes(c echo.Context) error {
	mode := c.QueryParam("mode")
	switch mode {
		case "random":
			return handleGetRandomNotes(c)
		default:
			defaultPageSize := 30
			page, err := strconv.Atoi(c.QueryParam("page"))
			fmt.Printf("\npage:\t%v\n", page)
			if err != nil {
				page = 1
			}
			if page < 1 {
				page = 1
			}
			pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
			fmt.Printf("\npage size:\t%v\n", pageSize)
			if err != nil {
				pageSize = defaultPageSize
			}
			if pageSize < 0 || pageSize > 100 {
				pageSize = defaultPageSize
			}
			notes, totalPages, err := db.ReadAllNotes(page, pageSize)
			if err != nil {
				return err
			}
			c.Response().Header().Set(echo.HeaderAccessControlExposeHeaders, "*")
			c.Response().Header().Set("Englishnote-Lastpage", strconv.Itoa(totalPages))
			return c.JSON(http.StatusOK, &notes)
	}
}

func handleGetRandomNotes(c echo.Context) error {
	defaultLimit := 30
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = defaultLimit
	}
	if limit < 1 {
		limit = defaultLimit
	}
	notes, err := db.GetRandomNotes(limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &notes)
}

func handleUpdateNotes(c echo.Context) error {
	notes := []Note{}
	if err := c.Bind(&notes); err != nil {
		return err
	}
	fmt.Printf("updating: %#v\n", notes)
	updatedNotes, err := db.UpdateNotes(notes)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &updatedNotes)
}

func handleDeleteNotes(c echo.Context) error {
	notes := []Note{}
	if err := c.Bind(&notes); err != nil {
		return err
	}
	err := db.DeleteNotes(notes)
	if err != nil {
		return err
	}
	for _, note := range notes {
		err = db.DeleteTags(note.Tags)
		if err != nil {
			return err
		}
	}
	return c.NoContent(http.StatusOK)
}

func handleGetRecord(c echo.Context) error {
	records, err := db.ReadAllRecords()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, records)
}

func handleAddRecord(c echo.Context) error {
	newRecords := []Record{}
	if err := c.Bind(&newRecords); err != nil {
		return err
	}
	records, err := db.AddRecords(newRecords)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &records)
}

