package value

type ObjectString struct {
	str string
}

func NewObjectString(str string) *ObjectString {
	return &ObjectString{str: str}
}

func (os *ObjectString) Type() ObjType {
	return ObjString
}

func (os *ObjectString) String() string {
	return os.str
}
