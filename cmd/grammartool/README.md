# grammartool

## Notes

**grammartool** is a basic front-end to the context-free grammar data
structure and associated functionality. It can provide basic information
about the grammar, including:

* Number of terminals, non-terminals and productions
* Whether the grammar is left-recursive
* Which non-terminals are nullable, or have cycles or e-productions

and (currently only for non-left-recursive grammars) can:

* verify whether the grammar recognizes a string
* verify whether the terminal regular expressions are valid
* generate Go code useful for including in a program
* show the output from the lexical analyzer

## Usage Examples

### Basic information

Specifying no options displays the grammar and some summary information:

	paul@horus:grammartool$ cat 1.grammar
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 1.grammar
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	The grammar is left-recursive.
	The grammar has 4 nonterminals, 5 terminals, and 7 productions
	paul@horus:grammartool$ cat 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	The grammar is not left-recursive.
	The grammar has 6 nonterminals, 5 terminals, and 9 productions
	paul@horus:grammartool$ 

### Detailed information

The `-a` option can give a variety of additional information about the
grammar:

	paul@horus:grammartool$ cat 7.grammar
	S : A B | C D
	A : E
	B : F
	C : E | e
	D : F | e
	E : `a`
	F : `b`
	paul@horus:grammartool$ grammartool -f 7.grammar -a
	S : A B | C D
	A : E
	B : F
	C : E | e
	D : F | e
	E : `a`
	F : `b`
	The grammar is not left-recursive.
	The grammar has 7 nonterminals, 2 terminals, and 10 productions
	The 7 nonterminals are S, A, B, C, D, E and F.
	The 2 terminals are `a` and `b`.
	No nonterminals have cycles.
	C and D have e-productions.
	S, C and D are nullable.
	No nonterminals are unreachable.
	No nonterminals are unproductive.
	First sets:
	First(S)  : { `a`, `b`, e }
	First(A)  : { `a` }
	First(B)  : { `b` }
	First(C)  : { `a`, e }
	First(D)  : { `b`, e }
	First(E)  : { `a` }
	First(F)  : { `b` }
	Follow sets:
	Follow(S) : { $ }
	Follow(A) : { `b` }
	Follow(B) : { $ }
	Follow(C) : { `b`, $ }
	Follow(D) : { $ }
	Follow(E) : { `b`, $ }
	Follow(F) : { $ }
	paul@horus:grammartool$ 

### Recognizing strings

The `-r` option shows whether the grammar recognizes a given string:

	paul@horus:grammartool$ cat 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 2.grammar -r -i='(3+4)*5'
	paul@horus:grammartool$ grammartool -f 2.grammar -r -i='(3+4)*5' -v
	Grammar recognizes string '(3+4)*5'.
	paul@horus:grammartool$ grammartool -f 2.grammar -r -i='(3+4)-5'
	Grammar does not recognize string '(3+4)-5'.
	paul@horus:grammartool$ 

### Generating a parse tree

The `-p` option generates a parse tree in bracketed format for a given
string:

	paul@horus:grammartool$ cat 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 2.grammar -p -i='(3+4)*5'
	(E (T (F `(` (E (T (F (Digits `3`)) (T' e)) (E' `+` (T (F (Digits `4`)) (T' e)) (E' e))) `)`) (T' `*` (F (Digits `5`)) (T' e))) (E' e))
	paul@horus:grammartool$ 

### Lexer output

The `-lex` option shows the output from the lexer for a given string:

	paul@horus:grammartool$ cat 1.grammar
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 1.grammar -lex -i='((3+4)*5)*(6+8)'
	T  n Lexeme
	-  - ------
	2  0 '('
	2  1 '('
	4  2 '3'
	0  3 '+'
	4  4 '4'
	3  5 ')'
	1  6 '*'
	4  7 '5'
	3  8 ')'
	1  9 '*'
	2 10 '('
	4 11 '6'
	0 12 '+'
	4 13 '8'
	3 14 ')'
	paul@horus:grammartool$

### Validating regular expressions

The `-checkRegexp` option validates that the terminal regular expressions
in the grammar compile with Go's regular expression engine:

	paul@horus:grammartool$ cat 98.grammar
	E      : E `\+` T | T
	T      : T `\*` F | F
	F      : `(` E `\)` | Digits
	Digits : `[[:digitals:]]+`
	paul@horus:grammartool$ grammartool -f 98.grammar -checkRegexp
	Regexp `(` did not compile: error parsing regexp: missing closing ): `(`
	Regexp `[[:digitals:]]+` did not compile: error parsing regexp: invalid character class range: `[:digitals:]`
	paul@horus:grammartool$ grammartool -f 98.grammar -checkRegexp -v
	Regexp `\+` compiled successfully.
	Regexp `\*` compiled successfully.
	Regexp `(` did not compile: error parsing regexp: missing closing ): `(`
	Regexp `\)` compiled successfully.
	Regexp `[[:digitals:]]+` did not compile: error parsing regexp: invalid character class range: `[:digitals:]`
	paul@horus:grammartool$ 

### Generating Go code

The `-generateSymbols` and `-generateGrammarString` options can be used to
output Go code declaring `const` identifiers for the terminals and
nonterminals in the grammar (which are useful for reading the parse tree)
and a string containing the grammar text (which is useful for embedding
a grammar in an executable):

	paul@horus:grammartool$ cat 2.grammar
	E      : T E'
	E'     : `\+` T E' | e
	T      : F T'
	T'     : `\*` F T' | e
	F      : `\(` E `\)` | Digits
	Digits : `[[:digit:]]+`
	paul@horus:grammartool$ grammartool -f 2.grammar -generateSymbols -generateGrammarString
	package main

	// Nonterminal identifier constants.
	const (
		ntE = iota // Nonterminal E
		ntT        // Nonterminal T
		ntEp       // Nonterminal E'
		ntF        // Nonterminal F
		ntTp       // Nonterminal T'
		ntDigits   // Nonterminal Digits
	)

	// Terminal identifier constants.
	const (
		t0 = iota // Terminal `\+`
		t1        // Terminal `\*`
		t2        // Terminal `\(`
		t3        // Terminal `\)`
		t4        // Terminal `[[:digit:]]+`
	)

	// grammarString is a string representation of the following grammar:
	// E      : T E'
	// E'     : `\+` T E' | e
	// T      : F T'
	// T'     : `\*` F T' | e
	// F      : `\(` E `\)` | Digits
	// Digits : `[[:digit:]]+`
	var grammarString = `E      : T E'
	E'     : ` + "`" + `\+` + "`" + ` T E' | e
	T      : F T'
	T'     : ` + "`" + `\*` + "`" + ` F T' | e
	F      : ` + "`" + `\(` + "`" + ` E ` + "`" + `\)` + "`" + ` | Digits
	Digits : ` + "`" + `[[:digit:]]+` + "`" + `
	`
	paul@horus:grammartool$ 

The `-nonTerminalPrefix`, `-terminalPrefix` and `-pkg` options may be used
to customize this generated output.
