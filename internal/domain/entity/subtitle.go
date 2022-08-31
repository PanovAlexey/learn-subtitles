package entity

import "time"

type Subtitle struct {
	Name      string
	Text      string
	CreatedAt time.Time
	Author    User
	IsDeleted bool
}
