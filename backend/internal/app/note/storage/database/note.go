package note

import (
	"fmt"
	"janstupica/StickyNotes/internal/app/models"

	supa "github.com/nedpals/supabase-go"
)

type NoteStorage struct {
	client    *supa.Client
	tableName string
}

func New(db *supa.Client) *NoteStorage {
	return &NoteStorage{
		client:    db,
		tableName: "note",
	}
}

func (n *NoteStorage) Create(note *models.Note) (*models.Note, error) {
	var result models.Note
	err := n.client.DB.From(n.tableName).Insert(note).Execute(&result)

	return &result, err
}

func (n *NoteStorage) Update(note *models.Note) (*models.Note, error) {
	var result models.Note
	err := n.client.DB.From(n.tableName).Update(note).Eq("id", fmt.Sprint(note.Id)).Execute(&result)

	return &result, err
}

func (n *NoteStorage) Delete(noteId int) error {
	err := n.client.DB.From(n.tableName).Delete().Eq("id", fmt.Sprint(noteId)).Execute(nil)

	return err
}

func (n *NoteStorage) Get(noteId int) (*models.Note, error) {
	var result models.Note
	err := n.client.DB.From(n.tableName).Select("*").Single().Eq("id", fmt.Sprint(noteId)).Execute(&result)

	return &result, err
}

func (n *NoteStorage) GetAll(userId int) (*[]models.Note, error) {
	var results []models.Note
	err := n.client.DB.From(n.tableName).Select("*").Eq("userId", fmt.Sprint(userId)).Execute(&results)

	return &results, err
}
