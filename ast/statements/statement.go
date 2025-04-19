package statements

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/jsonutil"
)

type Statement struct {
	Type    slang.StatementType `json:"-"`
	Comment string              `json:"-"`
	Core    Core                `json:"-"`
}

type Core interface {
	GetName() (string, bool)
	IsEqual(Core) bool
	Render(level int, w slang.CodeWriter)
}

func NewStatement(typ slang.StatementType, core Core) *Statement {
	return &Statement{
		Type:    typ,
		Comment: "",
		Core:    core,
	}
}

func (s *Statement) SetComment(comment string) {
	s.Comment = comment
}

func (s *Statement) GetComment() string {
	return s.Comment
}

func (s *Statement) GetType() slang.StatementType {
	return s.Type
}

func (s *Statement) GetName() (string, bool) {
	return s.Core.GetName()
}

func (s *Statement) Render(level int, w slang.CodeWriter) {
	if s.Comment != "" {
		for _, line := range strings.Split(s.Comment, "\n") {
			w.WriteIndent(level)
			w.WriteString("// ")
			w.WriteString(line)
			w.WriteNewline()
		}
	}

	w.WriteIndent(level)

	s.Core.Render(level, w)

	w.WriteNewline()
}

func (s *Statement) IsEqual(other slang.Statement) bool {
	s2, ok := other.(*Statement)
	if !ok {
		return false
	}

	if s.Type != s2.Type {
		return false
	}

	return s.Core.IsEqual(s2.Core)
}

func (s *Statement) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(s.Core)
	if err != nil {
		return []byte{}, err
	}

	d, err = sjson.SetBytes(d, "type", s.Type.String())
	if err != nil {
		return []byte{}, err
	}

	if s.Comment != "" {
		d, err = sjson.SetBytes(d, "comment", s.Comment)
		if err != nil {
			return []byte{}, err
		}
	}

	return d, nil
}

var (
	errTypeNotFound     = errors.New("type not found")
	errTypeNotString    = errors.New("type not a string")
	errCommentNotString = errors.New("comment not a string")
)

func (s *Statement) UnmarshalJSON(d []byte) error {
	result := gjson.GetBytes(d, "type")
	if !result.Exists() {
		return errTypeNotFound
	}

	if result.Type != gjson.String {
		return errTypeNotString
	}

	stmtType, ok := slang.ParseStatementTypeStr(result.String())
	if !ok {
		return fmt.Errorf("invalid statement type string '%s'", result.String())
	}

	var err error

	var core Core

	switch stmtType {
	case slang.StatementASSIGN:
		core, err = jsonutil.UnmarshalAs[Assign](d)
	case slang.StatementBREAK:
		core, err = jsonutil.UnmarshalAs[Break](d)
	case slang.StatementCOMMENT:
		core, err = jsonutil.UnmarshalAs[Comment](d)
	case slang.StatementCONST:
		core, err = jsonutil.UnmarshalAs[Const](d)
	case slang.StatementCONTINUE:
		core, err = jsonutil.UnmarshalAs[Const](d)
	case slang.StatementEXPRESSION:
		core, err = jsonutil.UnmarshalAs[Expression](d)
	case slang.StatementFOREACH:
		core, err = jsonutil.UnmarshalAs[ForEach](d)
	case slang.StatementFUNC:
		core, err = jsonutil.UnmarshalAs[Func](d)
	case slang.StatementIF:
		core, err = jsonutil.UnmarshalAs[If](d)
	case slang.StatementIFELSE:
		core, err = jsonutil.UnmarshalAs[IfElse](d)
	case slang.StatementRETURN:
		core, err = jsonutil.UnmarshalAs[Return](d)
	case slang.StatementSTRUCT:
		core, err = jsonutil.UnmarshalAs[Struct](d)
	case slang.StatementUSE:
		core, err = jsonutil.UnmarshalAs[Use](d)
	case slang.StatementVAR:
		core, err = jsonutil.UnmarshalAs[Var](d)
	}

	if err != nil {
		return fmt.Errorf("failed to unmarshal core: %w", err)
	}

	if core == nil {
		return fmt.Errorf("unhandled statement type %s", result.String())
	}

	s.Core = core
	s.Type = stmtType

	result = gjson.GetBytes(d, "comment")
	if result.Exists() {
		if result.Type != gjson.String {
			return errCommentNotString
		}

		s.Comment = result.String()
	}

	return nil
}

func statementsEqual(a, b *Statement) bool {
	return a.IsEqual(b)
}
