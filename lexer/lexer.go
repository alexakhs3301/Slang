package lexer

import "Goslang/token"

type Lexer struct {
	input        string //The input
	position     int    //Current position in input(points to current char)
	readPosition int    //Current reading position in input (after current char)
	ch           byte   //Current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '/':
		tok = newToken(token.SLASH, l.ch)

	case '*':
		tok = newToken(token.ASTERISK, l.ch)

	case '^':
		tok = newToken(token.POWER, l.ch)

	case '%':
		tok = newToken(token.MODULUS, l.ch)

	case '&':
		tok = newToken(token.BITAND, l.ch)

	case '|':
		tok = newToken(token.BITOR, l.ch)

	case '~':
		tok = newToken(token.BITNOT, l.ch)

	case '#':
		tok = newToken(token.BITXOR, l.ch)

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case ':':
		tok = newToken(token.COLON, l.ch)

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	case 'i':
		if l.peekChar() == 'n' && !l.isWhiteSpace(l.peekCharAt(1)) {
			// Check for 'int' keyword
			if l.peekCharAt(1) == 't' && (l.isWhiteSpace(l.peekCharAt(2)) || l.peekCharAt(2) == ';' || l.peekCharAt(2) == ',' || l.peekCharAt(2) == ')') {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal) // Assuming 'INT' is the token type for integers
				return tok
			}
		} else if l.peekChar() == 'f' && l.isWhiteSpace(l.peekCharAt(1)) {
			// Check for 'if' keyword
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // Assuming 'IF' is the token type for 'if' statements
			return tok
		}
		//treating 'i' and the rest as normal identifiers (IDENT Token)
		//tok = newToken(token.IDENT, l.ch)
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

	case 's':
		if l.peekChar() == 't' && !l.isWhiteSpace(l.peekCharAt(1)) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			//treating 's' as a normal identifier(IDENT Token)
			tok = newToken(token.IDENT, l.ch)
		}

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func (l *Lexer) isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekCharAt(offset int) byte {
	if l.readPosition+offset >= len(l.input) {
		return 0 // Represents end-of-file
	}
	return l.input[l.readPosition+offset]
}
