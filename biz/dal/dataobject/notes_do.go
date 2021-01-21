package dataobject

// NotesDO type
type NotesDO struct {
	Id						int64		`json:"id"`
	Title				 	string		`json:"title"`
	Content 		     	string		`json:"content"`
	Status					int8		`json:"status"`
	CreatedTime          	int32  		`db:"created_time"`
	UpdatedTime          	int32  		`db:"updated_time"`
}
