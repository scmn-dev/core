package app

import (
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/model"
)

// CreateNote creates a new note and saves it to the store
func CreateNote(s db.Store, dto *model.NoteDTO, schema string) (*model.Note, error) {
	rawModel := model.ToNote(dto)
	encModel := EncryptModel(rawModel)

	createdNote, err := s.Notes().Create(encModel.(*model.Note), schema)
	if err != nil {
		return nil, err
	}

	return createdNote, nil
}

// UpdateNote updates the note with the dto and applies the changes in the store
func UpdateNote(s db.Store, note *model.Note, dto *model.NoteDTO, schema string) (*model.Note, error) {
	rawModel := model.ToNote(dto)
	encModel := EncryptModel(rawModel).(*model.Note)

	note.Title = encModel.Title
	note.Note = encModel.Note

	updatedNote, err := s.Notes().Update(note, schema)
	if err != nil {
		return nil, err
	}

	return updatedNote, nil
}
