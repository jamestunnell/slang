package slang

type TokenSeq interface {
	Previous() *Token
	Current() *Token
	Next() *Token

	Advance()
	AdvanceSkip(skipTypes ...TokenType)
	Skip(skipTypes ...TokenType)
}
