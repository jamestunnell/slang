package parsers

import (
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type BodyParser interface {
	parsing.Parser

	GetStatements() []*statements.Statement
}

type bodyParser struct {
	*ParserBase

	Statements []*statements.Statement

	handleStartTok func(slang.TokenSeq) error
	endToken       slang.TokenType
	makeStmtParser MakeStmtParserFunc
}

type MakeStmtParserFunc func(cur *slang.Token) (StatementParser, error)

func NewBodyParser(
	handleStartTok func(slang.TokenSeq) error,
	endToken slang.TokenType,
	makeStmtParser MakeStmtParserFunc,
) *bodyParser {
	return &bodyParser{
		ParserBase:     NewParserBase(),
		Statements:     []*statements.Statement{},
		endToken:       endToken,
		handleStartTok: handleStartTok,
		makeStmtParser: makeStmtParser,
	}
}

func (p *bodyParser) GetStatements() []*statements.Statement {
	return p.Statements
}

func (p *bodyParser) readCommentLines(toks slang.TokenSeq) ([]string, bool) {
	var commentLines []string

	for toks.Current().Is(slang.TokenCOMMENT) {
		commentLines = append(commentLines, toks.Current().Value())

		toks.Advance()

		// always expect a newline after comment
		if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE) {
			return []string{}, false
		}

		toks.Advance()
	}

	return commentLines, true
}

func (p *bodyParser) addCommentStatement(lines []string) {
	stmt := statements.NewComment()

	stmt.SetComment(makeComment(lines))

	p.Statements = append(p.Statements, stmt)
}

func (p *bodyParser) Run(toks slang.TokenSeq) bool {
	p.Statements = []*statements.Statement{}

	if err := p.handleStartTok(toks); err != nil {
		p.errors = append(p.errors, parsing.NewParseError(err, toks.Current()))

		return false
	}

	_ = toks.Skip(slang.TokenNEWLINE)

	// empty block
	if toks.Current().Is(p.endToken) {
		toks.Advance()

		return true
	}

	for {
		var commentLines []string
		var ok bool

		commentLines, ok = p.readCommentLines(toks)
		if !ok {
			return false
		}

		if len(commentLines) > 0 {
			if toks.Current().Is(slang.TokenNEWLINE) {
				toks.Advance()

				p.addCommentStatement(commentLines)

				continue
			} else if toks.Current().Is(p.endToken) {
				toks.Advance()

				p.addCommentStatement(commentLines)

				return true
			}
		}

		if !p.parseStatement(toks, commentLines) {
			return false
		}

		numNewlines := toks.Skip(slang.TokenNEWLINE)

		if toks.Current().Is(p.endToken) {
			toks.Advance()

			break
		}

		// statements must be delimited with a newline
		if numNewlines == 0 {
			p.TokenErr(toks.Current(), slang.TokenNEWLINE)

			return false
		}
	}

	return true
}

func (p *bodyParser) parseStatement(toks slang.TokenSeq, commentLines []string) bool {
	stmtParser, err := p.makeStmtParser(toks.Current())
	if stmtParser == nil {
		p.errors = append(p.errors, parsing.NewParseError(err, toks.Current()))

		return false
	}

	if !p.RunSubStmtParser(toks, makeComment(commentLines), stmtParser) {
		return false
	}

	p.Statements = append(p.Statements, stmtParser.GetStatement())

	return true
}

func makeComment(lines []string) string {
	return strings.Join(lines, " ")
}
