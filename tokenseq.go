package slang

type TokenSeq interface {
	Previous() *Token
	Current() *Token
	Next() *Token

	Advance()
	AdvanceUntil(types ...TokenType) int
	AdvanceSkip(skipTypes ...TokenType) int
	Skip(skipTypes ...TokenType) int
}
