package dao

import (
	"context"
	"crud-note-simple/base"
	"crud-note-simple/biz/dal/dataobject"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type NotesDao struct {
	db *sqlx.DB
}

// NewAuthOpLogsDAO func
func NewNotesDao(db *sqlx.DB) *NotesDao {
	return &NotesDao{db}
}

func (dao *NotesDao) Insert(do *dataobject.NotesDO) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), base.SQLTimeOut)
	defer cancelFunc()

	var query = "insert into notes(title, content) values (:title, :content)"
	r, err := dao.db.NamedExecContext(myCtx, query, do)
	if err != nil {
		fmt.Printf("NamedExec in Insert(%v), error: %v", do, err)
		return 0
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Printf("LastInsertId in Insert(%v)_error: %v", do, err)
		return 0
	}
	return id
}

func (dao *NotesDao) Update(do *dataobject.NotesDO) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), base.SQLTimeOut)
	defer cancelFunc()

	var query = "UPDATE notes SET title = :title, content = :content WHERE  id = :id"
	r, err := dao.db.ExecContext(myCtx, query, do)
	if err != nil {
		fmt.Printf("ExecContext in Update(%v), error: %v", do, err)
		return 0
	}

	id, err := r.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected in Update(%v)_error: %v", do, err)
		return 0
	}
	return id
}

func (dao *NotesDao) DeleteById(noteId int64) int64 {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), base.SQLTimeOut)
	defer cancelFunc()

	var query = "UPDATE notes SET status = 1 WHERE id = :id"
	r, err := dao.db.ExecContext(myCtx, query, noteId)
	if err != nil {
		fmt.Printf("ExecContext in Update(_), error: %v", err)
		return 0
	}

	id, err := r.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected in Update(_)_error: %v", err)
		return 0
	}
	return id
}

func (dao *NotesDao) GetNoteById(noteId int64) *dataobject.NotesDO  {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), base.SQLTimeOut)
	defer cancelFunc()

	var query = "select * from notes where status = 0 and id = ?"
	r, err := dao.db.QueryxContext(myCtx, query, noteId)
	if err != nil {
		fmt.Printf("ExecContext in Update(_), error: %v", err)
		return nil
	}

	res := &dataobject.NotesDO{}
	if r.Next() {
		err = r.StructScan(res)
		if err != nil {
			fmt.Printf("StructScan in SelectChatsByQueryLink(_), error: %v", err)
		}
	} else {
		return nil
	}

	err = r.Err()
	if err != nil {
		fmt.Printf("rows in SelectChatsByQueryLink(_), error: %v", err)
		return nil
	}

	return res
}

func (dao *NotesDao) GetAllNote(offset, perPage int32) []*dataobject.NotesDO {
	myCtx, cancelFunc := context.WithTimeout(context.Background(), base.SQLTimeOut)
	defer cancelFunc()

	var query = "select * from notes where status = 0 order by create_time offset ? rows fetch next ? rows only"
	r, err := dao.db.QueryxContext(myCtx, query, offset, perPage)
	if err != nil {
		fmt.Printf("ExecContext in Update(_), error: %v", err)
		return []*dataobject.NotesDO{}
	}

	res := make([]*dataobject.NotesDO, 0)
	for r.Next() {
		do := &dataobject.NotesDO{}
		err = r.StructScan(do)
		if err != nil {
			fmt.Printf("StructScan in SelectChatsByQueryLink(_), error: %v", err)
			continue
		}
		res = append(res, do)
	}

	err = r.Err()
	if err != nil {
		fmt.Printf("rows in SelectChatsByQueryLink(_), error: %v", err)
		return []*dataobject.NotesDO{}
	}

	return res
}
