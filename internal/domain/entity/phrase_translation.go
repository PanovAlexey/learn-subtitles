package entity

import "time"

type PhraseTranslation struct {
	Text      string
	CreatedAt time.Time
	Author    User
	IsDeleted bool
}
