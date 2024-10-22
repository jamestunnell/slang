package parsing

import (
	"github.com/jamestunnell/slang/lexing"
)

type TokenSeq struct {
	lexer lexing.Lexer

	current, prev, next *lexing.Token
}

func NewTokenSeq(l lexing.Lexer) *TokenSeq {
	current := l.NextToken()
	next := l.NextToken()

	return &TokenSeq{
		lexer:   l,
		current: current,
		next:    next,
	}
}

func (seq *TokenSeq) Previous() *lexing.Token {
	return seq.prev
}

func (seq *TokenSeq) Current() *lexing.Token {
	return seq.current
}

// CurrIndex() int
func (seq *TokenSeq) Next() *lexing.Token {
	return seq.next
}

// func (seq *TokenSeq) Skip(types ...lexing.TokenType) {
// 	if slices.Contains(types, lexing.TokenEOF) {
// 		log.Fatal().Msg("cannot skip EOF")
// 	}

// 	for !slices.Contains(types, seq.tokens[seq.current].Info.Type) {
// 		seq.Advance()
// 	}
// }

func (seq *TokenSeq) Advance() {
	seq.prev = seq.current
	seq.current = seq.next
	seq.next = seq.lexer.NextToken()
}

func (seq *TokenSeq) AdvanceUntil(types ...lexing.TokenType) {
	types = append([]lexing.TokenType{lexing.TokenEOF}, types...)

	for !seq.current.Is(types...) {
		seq.Advance()
	}
}

func (seq *TokenSeq) AdvanceSkip(skipTypes ...lexing.TokenType) {
	seq.Advance()
	seq.Skip(skipTypes...)
}

func (seq *TokenSeq) Skip(skipTypes ...lexing.TokenType) {
	for seq.current.Is(skipTypes...) {
		seq.Advance()
	}
}
