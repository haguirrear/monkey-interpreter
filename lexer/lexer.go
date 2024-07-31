package lexer

import (
	"bytes"

	"github.com/haguirrear/monkey-interpreter/token"
)

type Lexer struct {
	input string

	// Current position in input (points to current char)
	position int
	// Current reading position in input (after current char)
	readPosition int
	// Current char under examination
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Reads one char and stores it in `ch` attribute
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWithespace()
	switch l.ch {
	case '=':
		ch := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewToken(token.EQUAL, bytes.NewBuffer([]byte{ch, l.ch}).String())
		} else {
			tok = token.NewTokenByte(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewTokenByte(token.PLUS, l.ch)
	case ',':
		tok = token.NewTokenByte(token.COMMA, l.ch)
	case ';':
		tok = token.NewTokenByte(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewTokenByte(token.LPAREN, l.ch)
	case ')':
		tok = token.NewTokenByte(token.RPAREN, l.ch)
	case '{':
		tok = token.NewTokenByte(token.LBRACE, l.ch)
	case '}':
		tok = token.NewTokenByte(token.RBRACE, l.ch)
	case '!':
		ch := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewToken(token.NOT_EQUAL, bytes.NewBuffer([]byte{ch, l.ch}).String())
		} else {
			tok = token.NewTokenByte(token.BANG, l.ch)
		}
	case '-':
		tok = token.NewTokenByte(token.MINUS, l.ch)
	case '/':
		tok = token.NewTokenByte(token.SLASH, l.ch)
	case '*':
		tok = token.NewTokenByte(token.ASTERISK, l.ch)
	case '<':
		tok = token.NewTokenByte(token.LOWER, l.ch)
	case '>':
		tok = token.NewTokenByte(token.GREATER, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readToken(isLetter)
			tok.Type = token.LookupIdentifier(tok.Literal)

			// returning because we already called readChar()
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readToken(isDigit)
			tok.Type = token.INT

			return tok
		}

		tok = token.NewTokenByte(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// reads a token literal while the func `isValid` is true
func (l *Lexer) readToken(isValid func(byte) bool) string {
	initialPos := l.position
	for isValid(l.ch) {
		l.readChar()
	}

	return l.input[initialPos:l.position]
}

func (l *Lexer) skipWithespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
