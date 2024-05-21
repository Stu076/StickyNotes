package note

import (
	"janstupica/StickyNotes/internal/app/models"
	"janstupica/StickyNotes/internal/app/note"
)

type NoteUseCase struct {
	storage note.Storage
}

func New(
	storage note.Storage,
) *NoteUseCase {
	return &NoteUseCase{
		storage: storage,
	}
}

func (n *NoteUseCase) Create(note *models.Note) (*[]models.Note, error) {
	result, err := n.storage.Create(note)

	return result, err
}

func (n *NoteUseCase) Update(note *models.Note) (*[]models.Note, error) {
	results, err := n.storage.Update(note)

	return results, err
}

func (n *NoteUseCase) Delete(noteId int) error {
	err := n.storage.Delete(noteId)

	return err
}

func (n *NoteUseCase) Get(noteId int) (*models.Note, error) {
	result, err := n.storage.Get(noteId)

	return result, err
}

func (n *NoteUseCase) GetAll(userId int) (*[]models.Note, error) {
	results, err := n.storage.GetAll(userId)

	return results, err
}
