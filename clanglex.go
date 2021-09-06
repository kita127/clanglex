package clanglex

// Lexicalize
func Lexicalize(src string) ([]*Token, error) {
	l := NewLexer(src)
	tokens, err := l.lexicalize()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
