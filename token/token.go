package token

type TokenType string

// const defines literal, which will just be repalced by its value in complilation
// these are some token types
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

// go has two ways of declaration
// either var a = ... or a := ...
// a := ... works in function
// var a = ... works everywhere
// you can also do var a int, without assignemnt
var lit2tkType = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func CheckIdent(ident string) TokenType {
	// go allows simple expression in if
	if tkType, ok := lit2tkType[ident]; ok {
		return tkType
	}
	return IDENT
}

func NewToken(ty TokenType, lit byte) Token {
	tok := Token{Type: ty, Literal: string(lit)}
	return tok
}
