package main

type Scanner interface {
	Current() rune
	Move() rune
	Peek() rune
	PeekTwo() rune
	Recale() rune // looks back
}

type StringScanner struct {
	source   []byte
	position int
}

func (s *StringScanner) Current() rune {
	if s.position >= len(s.source) {
		return -1
	}

	return rune(s.source[s.position])
}
func (s *StringScanner) Move() rune {
	s.position += 1

	if s.position >= len(s.source) {
		return -1
	}

	return rune(s.source[s.position])
}
func (s *StringScanner) Peek() rune {
	if s.position+1 >= len(s.source) {
		return -1
	}
	return rune(s.source[s.position+1])
}
func (s *StringScanner) PeekTwo() rune {
	if s.position+2 >= len(s.source) {
		return -1
	}
	return rune(s.source[s.position+2])
}
func (s *StringScanner) Recale() rune {
	if s.position < 1 {
		return -1
	}
	return rune(s.source[s.position-1])
}
