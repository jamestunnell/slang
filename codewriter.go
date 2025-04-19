package slang

import "strings"

type CodeWriter interface {
	WriteIndent(level int)
	WriteString(string)
	WriteNewline()

	String() string
}

type codeWriter struct {
	sb strings.Builder
}

func NewCodeWriter() *codeWriter {
	return &codeWriter{}
}

func (w *codeWriter) WriteIndent(level int) {
	for i := 0; i < level; i++ {
		w.sb.WriteRune('\t')
	}
}

func (w *codeWriter) WriteString(s string) {
	w.sb.WriteString(s)
}

func (w *codeWriter) WriteNewline() {
	w.sb.WriteRune('\n')
}

func (w *codeWriter) String() string {
	return w.sb.String()
}
