package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	POWER    = "^"
	MODULUS  = "%"

	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	//BITWISE OPERATORS
	BITAND = "&"
	BITOR  = "|"
	BITNOT = "~"
	BITXOR = "#"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUTH    = "TRUTH"
	LIE      = "LIE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	PRINT    = "PRINT"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"truth":  TRUTH,
	"lie":    LIE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"print":  PRINT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		switch ident {
		case "int":
			return INT
		case "string":
			return STRING
		default:
			return IDENT
		}
	}
}
