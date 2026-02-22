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
	EQ        = "=="
	NEQ       = "!="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	AST       = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LET       = "LET"
	FUNCTION  = "FUNCTION"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
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
var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func CheckIdent(ident string) TokenType {
	// go allows simple expression in if
	if tkType, ok := keywords[ident]; ok {
		return tkType
	}
	return IDENT
}

func NewToken(ty TokenType, lit byte) Token {
	tok := Token{Type: ty, Literal: string(lit)}
	return tok
}
