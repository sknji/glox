package value

type ObjType int

const (
	ObjString ObjType = iota + 1
	ObjFunction
)

type Object interface {
	Type() ObjType
	String() string
}
