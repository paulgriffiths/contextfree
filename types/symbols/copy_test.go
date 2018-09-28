package symbols

import "testing"

func TestStringCopy(t *testing.T) {
	s := String{NewNonTerminal(1), NewTerminal(2), NewSymbolEmpty()}
	c := s.Copy()
	c[0].I = 2
	c[1].I = 3
	if s[0].I != 1 {
		t.Errorf("nonterminal value changed, got %d, want %d", c[0].I, 1)
	}
	if s[1].I != 2 {
		t.Errorf("terminal value changed, got %d, want %d", c[1].I, 2)
	}
}

func TestStringListCopy(t *testing.T) {
	s := StringList{
		{NewNonTerminal(1), NewTerminal(2), NewSymbolEmpty()},
		{NewNonTerminal(3), NewTerminal(4), NewSymbolEmpty()},
	}
	c := s.Copy()
	c[0][0].I = 5
	c[1][1].I = 6
	if s[0][0].I != 1 {
		t.Errorf("nonterminal value changed, got %d, want %d", c[0][0].I, 1)
	}
	if s[1][1].I != 4 {
		t.Errorf("terminal value changed, got %d, want %d", c[1][1].I, 4)
	}
}
