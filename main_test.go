package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var testDBName = "test/test.db"

func TestMain(m *testing.M) {
	// before all test
	err := os.Remove(testDBName)
	if err != nil {
		os.Exit(1)
	}
	db, err = NewDB(testDBName)
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
	// test all test
}

func TestAddNote(t *testing.T) {
	fp, err := os.Open("test/new_note.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/note", fp)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleAddNote(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateNote(t *testing.T) {
	fp, err := os.Open("test/update_note.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/note", fp)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleUpdateNotes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteNote(t *testing.T) {
	fp, err := os.Open("test/delete_note.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/note", fp)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleDeleteNotes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAddRecord(t *testing.T) {
	fp, err := os.Open("test/new_record.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/record", fp)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleAddRecord(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

