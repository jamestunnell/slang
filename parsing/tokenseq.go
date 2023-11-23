package parsing

import (
	"github.com/jamestunnell/slang"
)

type TokenSeq struct {
	lexer slang.Lexer

	current, prev, next *slang.Token
}

func NewTokenSeq(l slang.Lexer) *TokenSeq {
	current := l.NextToken()
	next := l.NextToken()

	return &TokenSeq{
		lexer:   l,
		current: current,
		next:    next,
	}
}

func (seq *TokenSeq) Previous() *slang.Token {
	return seq.prev
}

func (seq *TokenSeq) Current() *slang.Token {
	return seq.current
}

// CurrIndex() int
func (seq *TokenSeq) Next() *slang.Token {
	return seq.next
}

// func (seq *TokenSeq) Skip(types ...slang.TokenType) {
// 	if slices.Contains(types, slang.TokenEOF) {
// 		log.Fatal().Msg("cannot skip EOF")
// 	}

// 	for !slices.Contains(types, seq.tokens[seq.current].Info.Type()) {
// 		seq.Advance()
// 	}
// }

func (seq *TokenSeq) Advance() {
	seq.prev = seq.current
	seq.current = seq.next
	seq.next = seq.lexer.NextToken()
}

func (seq *TokenSeq) AdvanceSkip(skipTypes ...slang.TokenType) {
	seq.Advance()
	seq.Skip(skipTypes...)
}

func (seq *TokenSeq) Skip(skipTypes ...slang.TokenType) {
	for seq.current.Is(skipTypes...) {
		seq.Advance()
	}
}
