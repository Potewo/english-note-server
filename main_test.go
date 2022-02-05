package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var testDBName = "testdata/test.db"
var cmpOpts = []cmp.Option {
	cmpopts.IgnoreFields(Note{}, "UpdatedAt", "CreatedAt", "DeletedAt"),
	cmpopts.IgnoreFields(Tag{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
	cmpopts.IgnoreFields(Record{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
}

func TestMain(m *testing.M) {
	// before all test
	_, err := os.Stat(testDBName)
	if err == nil {
		err := os.Remove(testDBName)
		if err != nil {
			os.Exit(1)
		}
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
	newNote, err := os.Open("testdata/new_note.json")
	if err != nil {
		t.Fatal(err)
	}
	defer newNote.Close()

	expectString, err := os.ReadFile("testdata/new_note_expect.json")
	if err != nil {
		t.Fatal(err)
	}
	expect := []Note{}
	json.Unmarshal(expectString, &expect)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/note", newNote)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleAddNote(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		recBody := []Note{}
		json.Unmarshal(rec.Body.Bytes(), &recBody)
		if len(expect) != len(recBody) {
			t.Errorf("length of records is different\n")
		}
		for i, got := range recBody {
			diff := cmp.Diff(&got, &expect[i], cmpOpts...)
			if len(diff) != 0 {
				t.Fatalf("differs: (-got +want)\n%s", diff)
			}
		}
	}
}

func TestUpdateNote(t *testing.T) {
	fp, err := os.Open("testdata/update_note.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	expectString, err := os.ReadFile("testdata/update_note_expect.json")
	if err != nil {
		t.Fatal(err)
	}
	expect := []Note{}
	json.Unmarshal(expectString, &expect)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/note", fp)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleUpdateNotes(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		recBody := []Note{}
		json.Unmarshal(rec.Body.Bytes(), &recBody)
		if len(expect) != len(recBody) {
			t.Errorf("length of records is different\n")
		}
		for i, got := range recBody {
			diff := cmp.Diff(&got, &expect[i], cmpOpts...)
			if len(diff) != 0 {
				t.Fatalf("differs: (-got +want)\n%s", diff)
			}
		}
	}
}

func TestDeleteNote(t *testing.T) {
	fp, err := os.Open("testdata/delete_note.json")
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

func TestGetNote(t *testing.T) {
	expectString, err := os.ReadFile("testdata/get_note_expect.json")
	if err != nil {
		t.Fatal(err)
	}
	expect := []Note{}
	json.Unmarshal(expectString, &expect)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/note", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleGetNotes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recBody := []Note{}
		json.Unmarshal(rec.Body.Bytes(), &recBody)
		if len(expect) != len(recBody) {
			t.Fatalf("length of notes is different\nexpect:\t%#v\ngot:\t%#v\n", expect, recBody)
		}
		t.Logf("%v", recBody)
		for i, got := range recBody {
			diff := cmp.Diff(&got, &expect[i], cmpOpts...)
			if len(diff) != 0 {
				t.Fatalf("differs: (-got +want)\n%s", diff)
			}
		}
	}
}

func TestAddRecord(t *testing.T) {
	newRecord, err := os.Open("testdata/new_record.json")
	if err != nil {
		t.Fatal(err)
	}
	defer newRecord.Close()

	expectString, err := os.ReadFile("testdata/new_record_expect.json")
	if err != nil {
		t.Fatal(err)
	}

	expect := []Record{}
	json.Unmarshal(expectString, &expect)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/record", newRecord)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleAddRecord(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		recBody := []Record{}
		json.Unmarshal(rec.Body.Bytes(), &recBody)
		if len(expect) != len(recBody) {
			t.Errorf("length of records is different\n")
		}
		for i, got := range recBody {
			diff := cmp.Diff(&got, &expect[i], cmpOpts...)
			if len(diff) != 0 {
				t.Fatalf("differs: (-got +want)\n%s", diff)
			}
		}
	}
}

func TestGetRecord(t *testing.T) {
	expectString, err := os.ReadFile("testdata/get_record_expect.json")
	if err != nil {
		t.Fatal(err)
	}
	expect := []Record{}
	json.Unmarshal(expectString, &expect)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/record", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, handleGetRecord(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		recBody := []Record{}
		json.Unmarshal(rec.Body.Bytes(), &recBody)
		if len(expect) != len(recBody) {
			t.Errorf("length of records is different\n")
		}
		for i, got := range recBody {
			diff := cmp.Diff(&got, &expect[i], cmpOpts...)
			if len(diff) != 0 {
				t.Fatalf("differs: (-got +want)\n%s", diff)
			}
		}
	}
}
