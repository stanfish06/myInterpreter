package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(ty TokenType, lit byte) Token {
	tok := Token{Type: ty, Literal: string(lit)}
	return tok
}
