package queryhelper

type QueryBody struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type URLWithID struct {
	ID     int `uri:"id"`
	NoteID int `uri:"noteId"`
}
