package syntax

type TokenStream struct {
	current int
	tokens  []*Token
}

func (s *TokenStream) Prev() *Token {
	s.current--
	return s.tokens[s.current]
}

func (s *TokenStream) Next() *Token {
	current := s.current
	s.current++
	return s.tokens[current]
}

func (s *TokenStream) PeekCurrent() *Token {
	return s.PeekAt(s.current)
}

func (s *TokenStream) Peek() *Token {
	return s.PeekAt(s.current + 1)
}

func (s *TokenStream) PeekAt(index int) *Token {
	if index >= len(s.tokens) {
		return s.tokens[len(s.tokens)-1]
	}

	return s.tokens[index]
}

func newTokenStream(tokens []*Token) *TokenStream {
	return &TokenStream{
		current: 0,
		tokens:  tokens,
	}
}
