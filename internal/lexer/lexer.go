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

func isLetter(ch byte) bool {
	// comparisons between char types are allowed (typical in c)
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

// In our first interpreter attempt, we used Consume 
// to truncate the string to keep track of our position
// This method avoids modifying the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position // save the first position
	for isLetter(l.ch) {
		l.readChar() // move the pointer to the last letter position
	}
	// Seems like we need a check here to see if a indentifer has been declared
	return l.input[startPosition:l.position] // return the slice of input that's the identifier
}

func (l *Lexer) readNumber() string {
	startPosition := l.position // save the first position
	for isDigit(l.ch) {
		l.readChar() // move the pointer to the last letter position
	}
	// Seems like we need a check here to see if a indentifer has been declared
	return l.input[startPosition:l.position] // return the slice of input that's the identifier
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		nextToken = newToken(token.ASSIGN, l.ch)
	case '+':
		nextToken = newToken(token.PLUS, l.ch)
	case '-':
		nextToken = newToken(token.MINUS, l.ch)
	case '!':
		nextToken = newToken(token.BANG, l.ch)
	case '*':
		nextToken = newToken(token.ASTERISK, l.ch)
	case '/':
		nextToken = newToken(token.SLASH, l.ch)
	case '<':
		nextToken = newToken(token.LT, l.ch)
	case '>':
		nextToken = newToken(token.GT, l.ch)
	case ',':
		nextToken = newToken(token.COMMA, l.ch)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		nextToken = newToken(token.LPAREN, l.ch)
	case ')':
		nextToken = newToken(token.RPAREN, l.ch)
	case '{':
		nextToken = newToken(token.LBRACE, l.ch)
	case '}':
		nextToken = newToken(token.RBRACE, l.ch)
	case 0:
		nextToken.Type = token.EOF
		nextToken.Literal = ""
	default:
		// Makes since that identifiers should be checked after
		// all known keywords and operators in the language
		if isLetter(l.ch) {
			nextToken.Literal = l.readIdentifier() // Works like readChar, so we'll return here
			nextToken.Type = token.LookupIdent(nextToken.Literal) // checks keyword types first. If none, returns IDENT
			return nextToken
		}
		if isDigit(l.ch) {
			nextToken.Literal = l.readNumber()
			nextToken.Type = token.INT
			return nextToken
		}
		// If something got here, it's a syntax error
		nextToken = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return nextToken
}
