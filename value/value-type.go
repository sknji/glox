package value

type ValueType int

const (
	ValBool ValueType = iota + 1
	ValNil
	ValNumber
	ValObj
)
