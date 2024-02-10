package slang

type TokenSeq interface {
	Previous() *Token
	Current() *Token
	Next() *Token

	Advance()
	AdvanceUntil(types ...TokenType)
	AdvanceSkip(skipTypes ...TokenType)
	Skip(skipTypes ...TokenType)
}
