package note_test

import (
	"janstupica/StickyNotes/config"
	"janstupica/StickyNotes/internal/app/database"
	"janstupica/StickyNotes/internal/app/models"
	note "janstupica/StickyNotes/internal/app/note/storage/database"
	"testing"
	"time"
)

var (
	db          *database.DB
	noteStorage *note.NoteStorage
	noteId      int
)

func setUp() {
	config.Init()
	db = database.New()
	noteStorage = note.New(db.Client)
}

func TestCreateNote(t *testing.T) {
	setUp()
	note := &models.Note{
		Id:        1,
		UserId:    1,
		Title:     "Testing Title",
		Content:   "Testing Content",
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	result, err := noteStorage.Create(note)

	if err != nil {
		t.Logf("Error: %s", err.Error())
		t.Fail()
	}

	if result != nil {

	} else {
		t.Logf("No result was returned")
		t.Fail()
	}
}
