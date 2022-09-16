package dto

// is needed to conveniently receive values passed by the user by pressing a button, rather than manually entering text.
type CommandButton struct {
	Data    string `json:"d"`
	Command string `json:"n"`
	Text    string `json:"t"`
}
