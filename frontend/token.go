package frontend

type TokenType int

const (
	// Single-character tokens.
	TokenLeftParen TokenType = iota
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

type Token struct {
	Type TokenType
	Val  string
	Line int
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
