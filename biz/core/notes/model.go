package notes

import (
	"crud-note-simple/biz/dal/dao"
	"crud-note-simple/biz/dal/dataobject"
	"github.com/jmoiron/sqlx"
)

type notesDao struct {
 	NotesDao *dao.NotesDao
}

type NotesModel struct {
	dao *notesDao
}

func NewNotesModel() *NotesModel {
	return &NotesModel{dao: &notesDao{}}
}

func(m *NotesModel) InstallModel(db *sqlx.DB){
	m.dao.NotesDao = dao.NewNotesDao(db)
}

func (m *NotesModel) CreateNote(title, content string) int64 {
	note := &dataobject.NotesDO{
		Title:       title,
		Content:     content,
	}

	id := m.dao.NotesDao.Insert(note)

	return id
}

func (m *NotesModel) UpdateNote(note *dataobject.NotesDO) bool {
	numRows := m.dao.NotesDao.Update(note)
	if numRows <= 0 {
		return false
	}
	return true
}

func (m *NotesModel) DeleteNote(noteId int64) bool {
	numRows := m.dao.NotesDao.DeleteById(noteId)
	if numRows <= 0 {
		return false
	}
	return true
}

func (m *NotesModel) GetNoteById(noteId int64) *dataobject.NotesDO {
	return m.dao.NotesDao.GetNoteById(noteId)
}

func (m *NotesModel) GetAllNotes(offsetPage, perPage int32) []*dataobject.NotesDO{
	offset := (offsetPage-1)*perPage

	return m.dao.NotesDao.GetAllNote(offset, perPage)
}
