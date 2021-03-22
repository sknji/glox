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
		TokenBang:         {c.unary, nil, precNone},
		TokenBangEqual:    {nil, c.binary, precNone},
		TokenEqual:        {nil, nil, precNone},
		TokenEqualEqual:   {nil, c.binary, precNone},
		TokenGreater:      {nil, c.binary, precNone},
		TokenGreaterEqual: {nil, c.binary, precNone},
		TokenLess:         {nil, c.binary, precNone},
		TokenLessEqual:    {nil, c.binary, precNone},
		TokenIdentifier:   {nil, nil, precNone},
		TokenString:       {nil, nil, precNone},
		TokenNumber:       {c.number, nil, precNone},
		TokenAnd:          {nil, nil, precNone},
		TokenClass:        {nil, nil, precNone},
		TokenElse:         {nil, nil, precNone},
		TokenFalse:        {c.literal, nil, precNone},
		TokenFor:          {nil, nil, precNone},
		TokenFun:          {nil, nil, precNone},
		TokenIf:           {nil, nil, precNone},
		TokenNil:          {c.literal, nil, precNone},
		TokenOr:           {nil, nil, precNone},
		TokenPrint:        {nil, nil, precNone},
		TokenReturn:       {nil, nil, precNone},
		TokenSuper:        {nil, nil, precNone},
		TokenThis:         {nil, nil, precNone},
		TokenTrue:         {c.literal, nil, precNone},
		TokenVar:          {nil, nil, precNone},
		TokenWhile:        {nil, nil, precNone},
		TokenError:        {nil, nil, precNone},
		TokenEof:          {nil, nil, precNone},
	}
}

func (c *Compiler) getRule(tokenType TokenType) *ParseRule {
	return c.parseRules[tokenType]
}
