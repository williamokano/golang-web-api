package entities

import "time"

type Login struct {
	Username       string
	HashedPassword string
	TOTPSecret     *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
