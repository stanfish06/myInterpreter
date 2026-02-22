package lexer

import tk "myInterpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is ASCII for NUL char
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Todo: handle int
func (l *Lexer) NextToken() tk.Token {
	var tok tk.Token
	l.skipWhitespace()
	nextChar := l.peekChar()
	switch l.ch {
	case '=':
		if nextChar == '=' {
			tok = tk.Token{Type: tk.EQ, Literal: string(l.ch) + string(nextChar)}
			l.readChar()
		} else {
			tok = tk.NewToken(tk.ASSIGN, l.ch)
		}
	case ';':
		tok = tk.NewToken(tk.SEMICOLON, l.ch)
	case '(':
		tok = tk.NewToken(tk.LPAREN, l.ch)
	case ')':
		tok = tk.NewToken(tk.RPAREN, l.ch)
	case ',':
		tok = tk.NewToken(tk.COMMA, l.ch)
	case '+':
		tok = tk.NewToken(tk.PLUS, l.ch)
	case '-':
		tok = tk.NewToken(tk.MINUS, l.ch)
	case '!':
		if nextChar == '=' {
			tok = tk.Token{Type: tk.NEQ, Literal: string(l.ch) + string(nextChar)}
			l.readChar()
		} else {
			tok = tk.NewToken(tk.BANG, l.ch)
		}
	case '*':
		tok = tk.NewToken(tk.AST, l.ch)
	case '/':
		tok = tk.NewToken(tk.SLASH, l.ch)
	case '<':
		tok = tk.NewToken(tk.LT, l.ch)
	case '>':
		tok = tk.NewToken(tk.GT, l.ch)
	case '{':
		tok = tk.NewToken(tk.LBRACE, l.ch)
	case '}':
		tok = tk.NewToken(tk.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tk.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = tk.CheckIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = tk.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = tk.NewToken(tk.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// used to read stuffs like let, fn, etc
func (l *Lexer) readIdentifier() string {
	// left position
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// char is rune in go, and can be convert to byte
// you can compare rune and byte
// rune -> int32 (also include other stuffs like unicodes)
// byte -> uint8
func isLetter(ch byte) bool {
	co1 := ch <= 'z' && ch >= 'a'
	co2 := ch <= 'Z' && ch >= 'A'
	co3 := ch == '_'
	return co1 || co2 || co3
}

// read number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch <= '9' && ch >= '0'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
