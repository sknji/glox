package value

type ObjType int

const (
	ObjString ObjType = iota + 1
)

type Object interface {
	Type() ObjType
	String() string
}
