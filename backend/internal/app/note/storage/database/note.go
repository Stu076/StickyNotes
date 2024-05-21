package note

import (
	"fmt"
	"janstupica/StickyNotes/internal/app/models"

	supa "github.com/nedpals/supabase-go"
)

type createNote struct {
	UserId  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type updateNote struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

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

func (n *NoteStorage) Create(note *models.Note) (*[]models.Note, error) {
	var result []models.Note
	newNote := &createNote{
		UserId:  note.UserId,
		Title:   note.Title,
		Content: note.Content,
	}

	err := n.client.DB.From(n.tableName).Insert(newNote).Execute(&result)

	return &result, err
}

func (n *NoteStorage) Update(note *models.Note) (*[]models.Note, error) {
	var result []models.Note
	newNote := &updateNote{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}

	err := n.client.DB.From(n.tableName).Update(newNote).Eq("id", fmt.Sprint(note.Id)).Execute(&result)

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
