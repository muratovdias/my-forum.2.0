package models

type Error struct {
	Status     int
	StatusText string
	Message    string
	Back       string
}
