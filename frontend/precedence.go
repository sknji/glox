package frontend

type Precedence uint

const (
	precNone Precedence = iota + 1
	precAssignment
	precOr
	precAnd
	precEquality
	precComparison
	precTerm
	precFactor
	precUnary
	precCall
	precPrimary
)

