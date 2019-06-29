package models

import "time"

type Faculty struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string
	LongName  string
}