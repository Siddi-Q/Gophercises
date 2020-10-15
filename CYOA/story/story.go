package story

// Story represents a Choose Your Own Adventure story.
// Each key is the name of a story chapter, and
// each value is a Chapter.
type Story map[string]Chapter

// Chapter represents a CYOA story chapter.
// Each chapter includes its title, the paragraphs it is composed
// of, and options available for the reader to take at the
// end of the chapter. If the options are empty it is
// assumed that you have reached the end of that particular
// story path.
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents a choice offered at the end of a story chapter.
// Text is the visible text end users will see,
// while the Chapter field will be the key to a chapter
// stored in the Story object this chapter was found in.
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
