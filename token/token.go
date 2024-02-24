package token

type TokenType string

type Token struct {
	Type 	TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	SLASH = "/"
	STAR = "*"
	MINUS = "-"
	BANG = "!"

	LT = "<"
	GT = ">"
	LEQT = "<="
	GEQT = ">="

	EQ = "=="
	DIFF = "!="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	DEF = "DEF"
	RETURN = "RETURN"
	IF = "IF"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"f": FUNCTION,
	"def": DEF,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}