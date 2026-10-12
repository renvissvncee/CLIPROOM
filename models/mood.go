package models

// Mood — настроение клипа. Закрытый словарь.
type Mood string

const (
	MoodHappy     Mood = "happy"
	MoodSad       Mood = "sad"
	MoodChill     Mood = "chill"
	MoodEnergetic Mood = "energetic"
	MoodDark      Mood = "dark"
)

// validMoods — приватный набор допустимых значений.
var validMoods = map[Mood]struct{}{
	MoodHappy:     {},
	MoodSad:       {},
	MoodChill:     {},
	MoodEnergetic: {},
	MoodDark:      {},
}

// IsValidMood — проверяет, что mood из закрытого словаря.
func (m Mood) IsValidMood() bool {
	_, ok := validMoods[m]
	return ok
}
