package internal

func NoComplexity() {
}

func OneIf() {
	if 1 == 2 {
	}
}

func And() {
	if 1 == 2 && 2 == 1 {
	}
}

func Or() {
	if 1 == 2 || 2 == 1 {
	}
}

type S struct {
}

func (s S) AFunction() {
}

func (s *S) PointerFunction() {
}
