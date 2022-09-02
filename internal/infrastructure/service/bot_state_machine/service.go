package bot_state_machine

import (
	"strconv"
)

var userDialogMap map[string]*Dialog

type UserStatesService struct {
	m map[string]*Dialog
}

func NewUserStatesService() UserStatesService {
	if userDialogMap == nil {
		userDialogMap = make(map[string]*Dialog)
	}

	return UserStatesService{
		m: userDialogMap,
	}
}

func (s UserStatesService) SetUserDialog(dialog *Dialog) {
	s.m[strconv.FormatInt(dialog.userId, 10)] = dialog
}

func (s UserStatesService) GetUserDialog(userId string) (*Dialog, bool) {
	dialog, ok := userDialogMap[userId]

	return dialog, ok
}
