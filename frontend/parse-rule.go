package frontend

type ParseFunc func()

type ParseRule struct {
	prefix     ParseFunc
	infix      ParseFunc
	precedence Precedence
}

func (c *Compiler) defineRules() {
	c.parseRules = map[TokenType]*ParseRule{
		TokenLeftParen:    {c.grouping, nil, precNone},
		TokenRightParen:   {nil, nil, precNone},
		TokenLeftBrace:    {nil, nil, precNone},
		TokenRightBrace:   {nil, nil, precNone},
		TokenComma:        {nil, nil, precNone},
		TokenDot:          {nil, nil, precNone},
		TokenMinus:        {c.unary, c.binary, precTerm},
		TokenPlus:         {nil, c.binary, precTerm},
		TokenSemicolon:    {nil, nil, precNone},
		TokenSlash:        {nil, c.binary, precFactor},
		TokenStar:         {nil, c.binary, precFactor},
		TokenBang:         {nil, nil, precNone},
		TokenBangEqual:    {nil, nil, precNone},
		TokenEqual:        {nil, nil, precNone},
		TokenEqualEqual:   {nil, nil, precNone},
		TokenGreater:      {nil, nil, precNone},
		TokenGreaterEqual: {nil, nil, precNone},
		TokenLess:         {nil, nil, precNone},
		TokenLessEqual:    {nil, nil, precNone},
		TokenIdentifier:   {nil, nil, precNone},
		TokenString:       {nil, nil, precNone},
		TokenNumber:       {c.number, nil, precNone},
		TokenAnd:          {nil, nil, precNone},
		TokenClass:        {nil, nil, precNone},
		TokenElse:         {nil, nil, precNone},
		TokenFalse:        {nil, nil, precNone},
		TokenFor:          {nil, nil, precNone},
		TokenFun:          {nil, nil, precNone},
		TokenIf:           {nil, nil, precNone},
		TokenNil:          {nil, nil, precNone},
		TokenOr:           {nil, nil, precNone},
		TokenPrint:        {nil, nil, precNone},
		TokenReturn:       {nil, nil, precNone},
		TokenSuper:        {nil, nil, precNone},
		TokenThis:         {nil, nil, precNone},
		TokenTrue:         {nil, nil, precNone},
		TokenVar:          {nil, nil, precNone},
		TokenWhile:        {nil, nil, precNone},
		TokenError:        {nil, nil, precNone},
		TokenEof:          {nil, nil, precNone},
	}
}

func (c *Compiler) getRule(tokenType TokenType) *ParseRule {
	return c.parseRules[tokenType]
}
