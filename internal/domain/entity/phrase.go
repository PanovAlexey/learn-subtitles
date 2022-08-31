package entity

import "time"

type Phrase struct {
	Text      string
	CreatedAt time.Time
	Author    User
	IsDeleted bool
}
