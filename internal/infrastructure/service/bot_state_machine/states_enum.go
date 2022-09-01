package bot_state_machine

type StatesEnum int

const (
	Rest StatesEnum = iota
	ReadyToAddSubtitlesName
	ReadyToAddSubtitlesText
	ReadyToAddSubtitlesProhibitedWords
	HasSubtitlesList
	SelectedSubtitles
	HasRandomPhrase
	ReadyToAddTranslationRandomPhrase
)
