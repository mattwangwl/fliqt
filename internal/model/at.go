package model

import "time"

type At struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
