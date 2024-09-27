package models

type Song struct {
	ID          string `json:"id" db:"id"`
	GroupName   string `json:"group_name" db:"group_name" binding:"required"`
	SongName    string `json:"song_name" db:"song_name"  binding:"required"`
	Link        string `json:"link" db:"link"`
	Text        string `json:"text" db:"text"`
	ReleaseDate string `json:"release_date" db:"release_date"`
}
