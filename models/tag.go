package models

// Tag — жанр/метка клипа. Закрытый словарь.
type Tag string

const (
	TagRock   Tag = "rock"
	TagPop    Tag = "pop"
	TagHipHop Tag = "hip-hop"
	TagJazz   Tag = "jazz"
	TagMetal  Tag = "metal"
)

var validTags = map[Tag]struct{}{
	TagRock:   {},
	TagPop:    {},
	TagHipHop: {},
	TagJazz:   {},
	TagMetal:  {},
}

// IsValidTag — проверяет, что tag из закрытого словаря.
func (t Tag) IsValidTag() bool {
	_, ok := validTags[t]
	return ok
}
