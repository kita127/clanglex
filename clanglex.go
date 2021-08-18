package clanglex

// Lexicalize
func Lexicalize(src string) []*Token {
	l := NewLexer(src)
	return l.lexicalize()
}
