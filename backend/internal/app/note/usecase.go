package note

import "janstupica/StickyNotes/internal/app/models"

type UseCase interface {
	Create(note *models.Note) (*[]models.Note, error)
	Update(note *models.Note) (*[]models.Note, error)
	Delete(noteId int) error
	Get(noteId int) (*models.Note, error)
	GetAll(userId int) (*[]models.Note, error)
}
