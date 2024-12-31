package lexing

type (
	And     struct{}
	Assign  struct{}
	Bang    struct{}
	Break   struct{}
	Class   struct{}
	Colon   struct{}
	Comma   struct{}
	Comment struct {
		val string
	}
	Const        struct{}
	Continue     struct{}
	DollarLBrace struct{}
	Dot          struct{}
	Else         struct{}
	Eof          struct{}
	Equal        struct{}
	False        struct{}
	Field        struct{}
	Float        struct {
		val string
	}
	ForEach      struct{}
	Func         struct{}
	Greater      struct{}
	GreaterEqual struct{}
	If           struct{}
	Illegal      struct{ val rune }
	In           struct{}
	Int          struct {
		val string
	}
	LBrace    struct{}
	LBracket  struct{}
	Less      struct{}
	LessEqual struct{}
	LParen    struct{}
	Method    struct{}
	Minus     struct{}
	Newline   struct{}
	NotEqual  struct{}
	Or        struct{}
	Plus      struct{}
	RBrace    struct{}
	RBracket  struct{}
	Return    struct{}
	RParen    struct{}
	Semicolon struct{}
	Slash     struct{}
	Star      struct{}
	String    struct{ val string }
	Symbol    struct {
		val string
	}
	True           struct{}
	Use            struct{}
	Var            struct{}
	VerbatimString struct{ val string }
)

const (
	StrAND          = "and"
	StrBANG         = "!"
	StrBREAK        = "break"
	StrCLASS        = "class"
	StrCOLON        = ":"
	StrCOMMA        = ","
	StrCONST        = "const"
	StrCONTINUE     = "continue"
	StrDOLLARBRACE  = "${"
	StrDOT          = "."
	StrELSE         = "else"
	StrEQUAL        = "="
	StrEQUALEQUAL   = "=="
	StrFALSE        = "false"
	StrFIELD        = "field"
	StrFOREACH      = "foreach"
	StrFUNC         = "func"
	StrGREATER      = ">"
	StrGREATEREQUAL = ">="
	StrIF           = "if"
	StrIN           = "in"
	StrLBRACE       = "{"
	StrLBRACKET     = "["
	StrLESS         = "<"
	StrLESSEQUAL    = "<="
	StrLPAREN       = "("
	StrMETHOD       = "method"
	StrMINUS        = "-"
	StrNEWLINE      = "\n"
	StrNOTEQUAL     = "!="
	StrOR           = "or"
	StrPLUS         = "+"
	StrRBRACE       = "}"
	StrRBRACKET     = "]"
	StrRETURN       = "return"
	StrRPAREN       = ")"
	StrSEMICOLON    = ";"
	StrSLASH        = "/"
	StrSTAR         = "*"
	StrTRUE         = "true"
	StrUSE          = "use"
	StrVAR          = "var"
)

func AND() *TokenInfo                      { return NewTokenInfo(TokenAND, StrAND) }
func BANG() *TokenInfo                     { return NewTokenInfo(TokenBANG, StrBANG) }
func BREAK() *TokenInfo                    { return NewTokenInfo(TokenBREAK, StrBREAK) }
func CLASS() *TokenInfo                    { return NewTokenInfo(TokenCLASS, StrCLASS) }
func COLON() *TokenInfo                    { return NewTokenInfo(TokenCOLON, StrCOLON) }
func COMMA() *TokenInfo                    { return NewTokenInfo(TokenCOMMA, StrCOMMA) }
func COMMENT(val string) *TokenInfo        { return NewTokenInfo(TokenCOMMENT, val) }
func CONST() *TokenInfo                    { return NewTokenInfo(TokenCONST, StrCONST) }
func CONTINUE() *TokenInfo                 { return NewTokenInfo(TokenCONTINUE, StrCONTINUE) }
func DOLLARLBRACE() *TokenInfo             { return NewTokenInfo(TokenDOLLARLBRACE, StrDOLLARBRACE) }
func DOT() *TokenInfo                      { return NewTokenInfo(TokenDOT, StrDOT) }
func ELSE() *TokenInfo                     { return NewTokenInfo(TokenELSE, StrELSE) }
func EOF() *TokenInfo                      { return NewTokenInfo(TokenEOF, "") }
func EQUAL() *TokenInfo                    { return NewTokenInfo(TokenEQUAL, StrEQUAL) }
func EQUALEQUAL() *TokenInfo               { return NewTokenInfo(TokenEQUALEQUAL, StrEQUALEQUAL) }
func FALSE() *TokenInfo                    { return NewTokenInfo(TokenFALSE, StrFALSE) }
func FIELD() *TokenInfo                    { return NewTokenInfo(TokenFIELD, StrFIELD) }
func FLOAT(val string) *TokenInfo          { return NewTokenInfo(TokenFLOAT, val) }
func FOREACH() *TokenInfo                  { return NewTokenInfo(TokenFOREACH, StrFOREACH) }
func FUNC() *TokenInfo                     { return NewTokenInfo(TokenFUNC, StrFUNC) }
func GREATER() *TokenInfo                  { return NewTokenInfo(TokenGREATER, StrGREATER) }
func GREATEREQUAL() *TokenInfo             { return NewTokenInfo(TokenGREATEREQUAL, StrGREATEREQUAL) }
func IF() *TokenInfo                       { return NewTokenInfo(TokenIF, StrIF) }
func ILLEGAL(val rune) *TokenInfo          { return NewTokenInfo(TokenILLEGAL, string([]rune{val})) }
func IN() *TokenInfo                       { return NewTokenInfo(TokenIN, StrIN) }
func INT(val string) *TokenInfo            { return NewTokenInfo(TokenINT, val) }
func LBRACE() *TokenInfo                   { return NewTokenInfo(TokenLBRACE, StrLBRACE) }
func LBRACKET() *TokenInfo                 { return NewTokenInfo(TokenLBRACKET, StrLBRACKET) }
func LESS() *TokenInfo                     { return NewTokenInfo(TokenLESS, StrLESS) }
func LESSEQUAL() *TokenInfo                { return NewTokenInfo(TokenLESSEQUAL, StrLESSEQUAL) }
func LPAREN() *TokenInfo                   { return NewTokenInfo(TokenLPAREN, StrLPAREN) }
func METHOD() *TokenInfo                   { return NewTokenInfo(TokenMETHOD, StrMETHOD) }
func MINUS() *TokenInfo                    { return NewTokenInfo(TokenMINUS, StrMINUS) }
func NEWLINE() *TokenInfo                  { return NewTokenInfo(TokenNEWLINE, StrNEWLINE) }
func NOTEQUAL() *TokenInfo                 { return NewTokenInfo(TokenNOTEQUAL, StrNOTEQUAL) }
func OR() *TokenInfo                       { return NewTokenInfo(TokenOR, StrOR) }
func PLUS() *TokenInfo                     { return NewTokenInfo(TokenPLUS, StrPLUS) }
func RBRACE() *TokenInfo                   { return NewTokenInfo(TokenRBRACE, StrRBRACE) }
func RBRACKET() *TokenInfo                 { return NewTokenInfo(TokenRBRACKET, StrRBRACKET) }
func RETURN() *TokenInfo                   { return NewTokenInfo(TokenRETURN, StrRETURN) }
func RPAREN() *TokenInfo                   { return NewTokenInfo(TokenRPAREN, StrRPAREN) }
func SEMICOLON() *TokenInfo                { return NewTokenInfo(TokenSEMICOLON, StrSEMICOLON) }
func SLASH() *TokenInfo                    { return NewTokenInfo(TokenSLASH, StrSLASH) }
func STAR() *TokenInfo                     { return NewTokenInfo(TokenSTAR, StrSTAR) }
func STRING(val string) *TokenInfo         { return NewTokenInfo(TokenSTRING, val) }
func SYMBOL(val string) *TokenInfo         { return NewTokenInfo(TokenSYMBOL, val) }
func TRUE() *TokenInfo                     { return NewTokenInfo(TokenTRUE, StrTRUE) }
func USE() *TokenInfo                      { return NewTokenInfo(TokenUSE, StrUSE) }
func VAR() *TokenInfo                      { return NewTokenInfo(TokenVAR, StrVAR) }
func VERBATIMSTRING(val string) *TokenInfo { return NewTokenInfo(TokenVERBATIMSTRING, val) }
