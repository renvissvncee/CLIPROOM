package models

type Clip struct {
	id          int
	title       string
	artist      string
	description string

	likes uint

	tags  []string
	moods []string

	link string
}
