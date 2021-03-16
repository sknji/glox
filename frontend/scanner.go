package frontend

type Scanner struct {
	current int
	line    int
	source  []byte
	length  int
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		current: 0,
		line:    1,
		source:  source,
		length:  len(source),
	}
}

// Each call scans a complete token. We know we are at the beginning of
// a new token when we enter the function.
func (s *Scanner) scanToken() *Token {
	s.skipWhitespace()

	if s.isAtEnd() {
		return s.makeToken("", TokenEof)
	}

	c := s.advance()
	if isAlpha(c) {
		return s.identifier()
	}

	if isDigit(c) {
		return s.number()
	}

	switch c {
	case '(':
		return s.makeToken("(", TokenLeftParen)
	case ')':
		return s.makeToken(")", TokenRightParen)
	case '{':
		return s.makeToken("{", TokenLeftBrace)
	case '}':
		return s.makeToken("}", TokenRightBrace)
	case ';':
		return s.makeToken(";", TokenSemicolon)
	case ',':
		return s.makeToken(",", TokenComma)
	case '.':
		return s.makeToken(".", TokenDot)
	case '-':
		return s.makeToken("-", TokenMinus)
	case '+':
		return s.makeToken("+", TokenPlus)
	case '/':
		if s.peekNext() == '/' {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			return s.makeToken("/", TokenSlash)
		}
	case '*':
		return s.makeToken("*", TokenStar)
	case '!':
		if s.match('=') {
			return s.makeToken("!=", TokenBangEqual)
		} else {
			return s.makeToken("!", TokenBang)
		}
	case '=':
		if s.match('=') {
			return s.makeToken("==", TokenEqualEqual)
		} else {
			return s.makeToken("=", TokenEqual)
		}
	case '<':
		if s.match('=') {
			return s.makeToken("<=", TokenLessEqual)
		} else {
			return s.makeToken("<", TokenLess)
		}
	case '>':
		if s.match('=') {
			return s.makeToken(">=", TokenGreaterEqual)
		} else {
			return s.makeToken(">", TokenGreater)
		}
	case '\n':
		s.line += 1
		s.advance()
	case '"':
		return s.string()

	}

	return s.errorToken("Unexpected character.")
}

func (s *Scanner) makeToken(val string, _type TokenType) *Token {
	return &Token{
		Type: _type,
		Line: s.line,
		Val:  val,
	}
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) errorToken(message string) *Token {
	return &Token{
		Type: TokenError,
		Line: s.line,
		Val:  message,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= s.length
}

func (s *Scanner) advance() byte {
	s.current += 1
	return s.source[s.current]
}

// Handle all the white spaces apart from the new lines
func (s *Scanner) skipWhitespace() {
	for {
		switch s.peek() {
		case ' ', '\r', '\t':
			s.advance()
		default:
			return
		}
	}
}

func (s *Scanner) peek() byte {
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current+1]
}

func (s *Scanner) string() *Token {
	start := s.current
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line += 1
		}

		s.advance()
	}

	if s.isAtEnd() {
		return s.errorToken("Unterminated string.")
	}

	// Consume the closing quote
	s.advance()

	return s.makeToken(string(s.source[start:s.current]), TokenString)
}

func (s *Scanner) number() *Token {
	start := s.current
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
	}

	for isDigit(s.peek()) {
		s.advance()
	}

	return s.makeToken(string(s.source[start:s.current]), TokenNumber)
}

func (s *Scanner) identifier() *Token {
	start := s.current
	for isAlpha(s.peek()) || isDigit(s.peek()) {
		s.advance()
	}

	ident := string(s.source[start:s.current])
	return s.makeToken(ident, s.identifierType(ident))
}

func (s *Scanner) identifierType(ident string) TokenType {
	if tokenType, ok := Keywords[ident]; ok {
		return tokenType
	}

	return TokenIdentifier
}
