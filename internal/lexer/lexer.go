package lexer

import (
	"github.com/nicolasjhampton/simple_interpreter/internal/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	switch l.ch {
	case '=':
		nextToken = newToken(token.ASSIGN, l.ch)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		nextToken = newToken(token.LPAREN, l.ch)
	case ')':
		nextToken = newToken(token.RPAREN, l.ch)
	case ',':
		nextToken = newToken(token.COMMA, l.ch)
	case '+':
		nextToken = newToken(token.PLUS, l.ch)
	case '{':
		nextToken = newToken(token.LBRACE, l.ch)
	case '}':
		nextToken = newToken(token.RBRACE, l.ch)
	case 0:
		nextToken.Type = token.EOF
		nextToken.Literal = ""
	}

	l.readChar()
	return nextToken
}
