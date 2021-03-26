package frontend

import (
	"errors"
	"math"
)

const localScopeLimit = math.MaxUint8

var localScopeLimitError = errors.New("too many local variables in function")

type CompilerScope struct {
	locals     [localScopeLimit]*Local
	localCount int
	scopeDepth int
}

type Local struct {
	token *Token
	depth int
}

func NewCompilerScope() *CompilerScope {
	return &CompilerScope{
		locals:     [localScopeLimit]*Local{},
		localCount: 0,
		scopeDepth: 0,
	}
}

func (cs *CompilerScope) addLocal(tok *Token) error {
	if cs.localCount == localScopeLimit {
		return localScopeLimitError
	}
	local := &Local{
		token: tok,
		depth: -1,
	}
	cs.locals[cs.localCount] = local
	cs.localCount += 1

	return nil
}
