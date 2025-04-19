package field

import (
	"strings"

	"github.com/jamestunnell/slang"
)

type Seq []*Field

type namesType struct {
	Names []string
	Type  slang.Type
}

func (fields Seq) makeNamesTypes() []namesType {
	nts := []namesType{}

	if len(fields) == 0 {
		return nts
	}

	names := []string{fields[0].Name}
	typ := fields[0].Type

	for i := 1; i < len(fields); i++ {
		if fields[i].Type.IsEqual(typ) {
			names = append(names, fields[i].Name)

			continue
		}

		nts = append(nts, namesType{names, typ})
		names = []string{fields[i].Name}
		typ = fields[i].Type
	}

	nts = append(nts, namesType{names, typ})

	return nts
}

func (fields Seq) Render(level int, w slang.CodeWriter) {
	nts := fields.makeNamesTypes()

	switch len(nts) {
	case 0:
		w.WriteString("()")
	case 1:
		w.WriteString("(")

		nts[0].Render(w)

		w.WriteString(")")
	default:
		w.WriteString("(")

		for _, nt := range nts {
			w.WriteNewline()
			w.WriteIndent(level + 1)

			nt.Render(w)
		}

		w.WriteString(")")
	}
}

func (nt namesType) Render(w slang.CodeWriter) {
	w.WriteString(strings.Join(nt.Names, ", "))
	w.WriteString(" ")
	w.WriteString(nt.Type.String())
}
