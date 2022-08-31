package entity

import "time"

type User struct {
	FirstName string
	LastName  string
	Login     string
	CreatedAt time.Time
	IsDeleted bool
}
