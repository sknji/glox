package frontend

type TokenType int

const (
	// Single-character tokens.
	TokenLeftParen TokenType = iota + 1
	TokenRightParen
	TokenLeftBrace
	TokenRightBrace
	TokenComma
	TokenDot
	TokenMinus
	TokenPlus
	TokenSemicolon
	TokenSlash
	TokenStar

	// One or two character tokens.
	TokenBang
	TokenBangEqual
	TokenEqual
	TokenEqualEqual
	TokenGreater
	TokenGreaterEqual
	TokenLess
	TokenLessEqual

	// Literals.
	TokenIdentifier
	TokenString
	TokenNumber

	// Keywords.
	TokenAnd
	TokenClass
	TokenElse
	TokenFalse
	TokenFor
	TokenFun
	TokenIf
	TokenNil
	TokenOr
	TokenPrint
	TokenReturn
	TokenSuper
	TokenThis
	TokenTrue
	TokenVar
	TokenWhile

	TokenError
	TokenEof
)

func (tt TokenType) String() string {
	return getTokenStr(tt)
}

func (tt TokenType) Is(tokenType TokenType) bool {
	return tt == tokenType
}

type Token struct {
	Type TokenType
	Val  string
	Line uint
}

var Keywords = map[string]TokenType{
	"and":    TokenAnd,
	"class":  TokenClass,
	"ese":    TokenElse,
	"if":     TokenIf,
	"nil":    TokenNil,
	"or":     TokenOr,
	"print":  TokenPrint,
	"return": TokenReturn,
	"super":  TokenSuper,
	"var":    TokenVar,
	"while":  TokenWhile,
	"false":  TokenFalse,
	"for":    TokenFor,
	"fun":    TokenFun,
	"this":   TokenThis,
	"true":   TokenTrue,
}

var tokenTypeToString = map[TokenType]string{
	TokenLeftParen:  "TokenLeftParen",
	TokenRightParen: "TokenRightParen",
	TokenLeftBrace:  "TokenLeftBrace",
	TokenRightBrace: "TokenRightBrace",
	TokenComma:      "TokenComma",
	TokenDot:        "TokenDot",
	TokenMinus:      "TokenMinus",
	TokenPlus:       "TokenPlus",
	TokenSemicolon:  "TokenSemicolon",
	TokenSlash:      "TokenSlash",
	TokenStar:       "TokenStar",

	// One or two character tokens.
	TokenBang:         "TokenBang",
	TokenBangEqual:    "TokenBangEqual",
	TokenEqual:        "TokenEqual",
	TokenEqualEqual:   "TokenEqualEqual",
	TokenGreater:      "TokenGreater",
	TokenGreaterEqual: "TokenGreaterEqual",
	TokenLess:         "TokenLess",
	TokenLessEqual:    "TokenLessEqual",

	// Literals.
	TokenIdentifier: "TokenIdentifier",
	TokenString:     "TokenString",
	TokenNumber:     "TokenNumber",

	// Keywords.
	TokenAnd:    "TokenAnd",
	TokenClass:  "TokenClass",
	TokenElse:   "TokenElse",
	TokenFalse:  "TokenFalse",
	TokenFor:    "TokenFor",
	TokenFun:    "TokenFun",
	TokenIf:     "TokenIf",
	TokenNil:    "TokenNil",
	TokenOr:     "TokenOr",
	TokenPrint:  "TokenPrint",
	TokenReturn: "TokenReturn",
	TokenSuper:  "TokenSuper",
	TokenThis:   "TokenThis",
	TokenTrue:   "TokenTrue",
	TokenVar:    "TokenVar",
	TokenWhile:  "TokenWhile",

	TokenError: "TokenError",
	TokenEof:   "TokenEof",
}

func getTokenStr(tokType TokenType) string {
	return tokenTypeToString[tokType]
}
