package parsing

import (
	"github.com/jamestunnell/slang"
)

type TokenSeq interface {
	Current() *slang.Token
	Next() *slang.Token

	Advance()
	AdvanceUntil(types ...slang.TokenType) int
	AdvanceSkip(skipTypes ...slang.TokenType) int
	Skip(skipTypes ...slang.TokenType) int
}

type tokenSeq struct {
	lexer slang.Lexer

	current, next *slang.Token
}

func NewTokenSeq(l slang.Lexer) TokenSeq {
	current := l.NextToken()
	next := l.NextToken()

	return &tokenSeq{
		lexer:   l,
		current: current,
		next:    next,
	}
}

func (seq *tokenSeq) Current() *slang.Token {
	return seq.current
}

// CurrIndex() int
func (seq *tokenSeq) Next() *slang.Token {
	return seq.next
}

// func (seq *tokenSeq) Skip(types ...slang.TokenType) {
// 	if slices.Contains(types, slang.TokenEOF) {
// 		log.Fatal().Msg("cannot skip EOF")
// 	}

// 	for !slices.Contains(types, seq.tokens[seq.current].Info.Type()) {
// 		seq.Advance()
// 	}
// }

func (seq *tokenSeq) Advance() {
	seq.current = seq.next
	seq.next = seq.lexer.NextToken()
}

func (seq *tokenSeq) AdvanceUntil(types ...slang.TokenType) int {
	types = append([]slang.TokenType{slang.TokenEOF}, types...)

	advances := 0

	for !seq.current.Is(types...) {
		seq.Advance()

		advances++
	}

	return advances
}

func (seq *tokenSeq) AdvanceSkip(skipTypes ...slang.TokenType) int {
	seq.Advance()

	return seq.Skip(skipTypes...)
}

func (seq *tokenSeq) Skip(skipTypes ...slang.TokenType) int {
	advances := 0

	for seq.current.Is(skipTypes...) {
		seq.Advance()

		advances++
	}

	return advances
}
