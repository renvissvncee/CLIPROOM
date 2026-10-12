package models

type Clip struct {
	ID          uint
	Title       string
	Artist      string
	Description string

	Likes uint

	Tags  []Tag
	Moods []Mood

	Link string
}
