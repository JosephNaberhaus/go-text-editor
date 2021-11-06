package editor

const (
	DefaultWidth = 80
)

type TextEditor struct {
	// The number of characters to indent the first line
	firstLineIndent int

	// The number of grapheme clusters that can be placed on one line
	width int

	// The text content of the text editor
	paragraphs []paragraph

	// Stores the paragraph that the user's cursor is on
	cursorParagraph int

	// Stores the position of the cursor
	cursorPos int

	// The preferred column that the cursor will try to navigate to when the up/down navigations are used.
	cursorPreferredColumn int

	wrappedLinesCache [][]paragraph
}

// NewEditor create a new empty text editor.
func NewEditor() *TextEditor {
	return &TextEditor{
		width: DefaultWidth,
		paragraphs: []paragraph{
			make([]graphemeCluster, 0),
		},
	}
}

func (t *TextEditor) SetFirstLineIndent(newFirstLineIndent int) {
	t.firstLineIndent = newFirstLineIndent
	t.cursorPos = t.minCursorPos()
	t.cursorPreferredColumn = 0
	t.wrappedLinesCache = nil
}

func (t *TextEditor) SetWidth(newWidth int) {
	t.width = newWidth
	t.cursorPos = t.minCursorPos()
	t.cursorPreferredColumn = 0
	t.wrappedLinesCache = nil
}
