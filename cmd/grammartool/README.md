# grammar_info

## Notes

**grammar_info** is a basic front-end to the context-free grammar data
structure and associated functionality. It can provide basic information
about the grammar, including:

* Number of terminals, non-terminals and productions
* Whether the grammar is left-recursive
* Which non-terminals are nullable, or have cycles or e-productions

and (currently only for non-left-recursive grammars) can verify whether the
grammar recognizes a string.

## Usage Examples

A sample session:

	paul@horus:grammar_info$ cat 1.grammar
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammar_info$ ./grammar_info -f 1.grammar -a
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`

	The grammar is left-recursive.
	The grammar has 4 nonterminals, 5 terminals, and 7 productions
	The 4 nonterminals are E, T, F and Digits.
	The 5 terminals are `\+`, `\*`, `\(`, `\)` and `[[:digit:]]+`.
	No nonterminals have cycles.
	No nonterminals have e-productions.
	No nonterminals are nullable.
	paul@horus:grammar_info$ cat 6.grammar
	S : A | `a` | `b`
	A : S | `c` | `d`
	paul@horus:grammar_info$ ./grammar_info -f 6.grammar -a
	S : A | `a` | `b`
	A : S | `c` | `d`

	The grammar is left-recursive.
	The grammar has 2 nonterminals, 4 terminals, and 6 productions
	The 2 nonterminals are S and A.
	The 4 terminals are `a`, `b`, `c` and `d`.
	S and A have cycles.
	No nonterminals have e-productions.
	No nonterminals are nullable.
	paul@horus:grammar_info$ cat 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammar_info$ ./grammar_info -f 2.grammar -a
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`

	The grammar is not left-recursive.
	The grammar has 6 nonterminals, 5 terminals, and 9 productions
	The 6 nonterminals are E, T, E', F, T' and Digits.
	The 5 terminals are `\+`, `\*`, `\(`, `\)` and `[[:digit:]]+`.
	No nonterminals have cycles.
	E' and T' have e-productions.
	E' and T' are nullable.
	paul@horus:grammar_info$ ./grammar_info -f 2.grammar -r '(3+4)*5'
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`

	The grammar is not left-recursive.
	The grammar has 6 nonterminals, 5 terminals, and 9 productions
	Grammar recognizes string '(3+4)*5'.
	paul@horus:grammar_info$ ./grammar_info -f 2.grammar -g=false -s=false -r '(3+4)/5'
	Grammar does not recognize string '(3+4)/5'.
	paul@horus:grammar_info$ cat 11.grammar
	S : `0` S `1` | `01`
	paul@horus:grammar_info$ ./grammar_info -f 11.grammar -p '000111'
	S : `0` S `1` | `01`

	The grammar is not left-recursive.
	The grammar has 1 nonterminal, 3 terminals, and 2 productions
	Parse tree for string '000111': (S `0` (S `0` (S `01`) `1`) `1`)
	paul@horus:grammar_info$ 

