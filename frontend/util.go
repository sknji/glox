package frontend

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		ch == '_'
}
