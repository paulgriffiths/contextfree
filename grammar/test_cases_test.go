package grammar_test

type grammarTestCase struct {
	filename                 string
	isLeftRecursive          bool
	numNonTerminals          int
	numTerminals             int
	numProductions           int
	nonTerminalNames         []string
	terminalNames            []string
	leftRecursive            []string
	immediatelyLeftRecursive []string
	haveCycles               []string
	haveEProds               []string
	areNullable              []string
	unreachable              []string
	unproductive             []string
	firsts                   map[string][]string
	follows                  map[string][]string
}

var grammarTestCases = []grammarTestCase{
	{
		tgArithLr, true, 3, 5, 6,
		[]string{"E", "T", "F"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{"E", "T"}, []string{"E", "T"},
		[]string{}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"E": []string{"\\(", "[[:digit:]]+"},
			"T": []string{"\\(", "[[:digit:]]+"},
			"F": []string{"\\(", "[[:digit:]]+"},
			//"Digits": []string{"[[:digit:]]+"},
		},
		map[string][]string{
			"F": []string{"\\+", "\\*", "\\)", "$"},
			"T": []string{"\\+", "\\*", "\\)", "$"},
			"E": []string{"\\+", "\\)", "$"},
			//"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgArithNlr, false, 6, 5, 9,
		[]string{"E", "T", "E'", "F", "T'", "Digits"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{}, []string{},
		[]string{}, []string{"E'", "T'"}, []string{"E'", "T'"},
		[]string{},
		[]string{},
		map[string][]string{
			"E":      []string{"\\(", "[[:digit:]]+"},
			"T":      []string{"\\(", "[[:digit:]]+"},
			"E'":     []string{"\\+", ""},
			"F":      []string{"\\(", "[[:digit:]]+"},
			"T'":     []string{"\\*", ""},
			"Digits": []string{"[[:digit:]]+"},
		},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\)", "$"},
			"E":      []string{"\\)", "$"},
			"E'":     []string{"\\)", "$"},
			"T'":     []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgArithAmbig, true, 2, 5, 5,
		[]string{"E", "Digits"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{"E"}, []string{"E"},
		[]string{}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"E":      []string{"\\(", "[[:digit:]]+"},
			"Digits": []string{"[[:digit:]]+"},
		},
		map[string][]string{
			"E":      []string{"\\*", "\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
		},
	},
	{
		tgBalParens1, true, 1, 2, 2,
		[]string{"S"},
		[]string{"\\(", "\\)"},
		[]string{"S"}, []string{"S"},
		[]string{}, []string{"S"}, []string{"S"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"\\(", ""},
		},
		map[string][]string{
			"S": []string{"\\(", "\\)", "$"},
		},
	},
	{
		tgBalParens2, false, 1, 2, 2,
		[]string{"S"},
		[]string{"\\(", "\\)"},
		[]string{}, []string{},
		[]string{}, []string{"S"}, []string{"S"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"\\(", ""},
		},
		map[string][]string{
			"S": []string{"\\)", "$"},
		},
	},
	{
		tgZeroOne, false, 1, 3, 2,
		[]string{"S"},
		[]string{"0", "1", "01"},
		[]string{}, []string{},
		[]string{}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"0", "01"},
		},
		map[string][]string{
			"S": []string{"1", "$"},
		},
	},
	{
		tgIndirectLr1, true, 2, 4, 5,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"S", "A"}, []string{"A"},
		[]string{}, []string{"A"}, []string{"A"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c"},
			"A": []string{"a", "b", "c", ""},
		},
		map[string][]string{
			"S": []string{"d", "$"},
			"A": []string{"a", "c"},
		},
	},
	{
		tgIndirectLr2, true, 4, 5, 8,
		[]string{"S", "A", "B", "C"},
		[]string{"a", "b", "c", "d", "e"},
		[]string{"S", "A", "B", "C"}, []string{},
		[]string{}, []string{"A", "B", "C"}, []string{"A", "B", "C"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c", "d"},
			"A": []string{"a", "b", "c", "d", ""},
			"B": []string{"a", "b", "c", "d", ""},
			"C": []string{"a", "b", "c", "d", ""},
		},
		map[string][]string{
			"S": []string{"e", "$"},
			"A": []string{"a"},
			"B": []string{"c"},
			"C": []string{"d"},
		},
	},
	{
		tgIndirectLr3, true, 5, 6, 10,
		[]string{"S", "A", "B", "C", "D"},
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"A", "B", "C", "D"}, []string{},
		[]string{},
		[]string{"A", "B", "C", "D"}, []string{"A", "B", "C", "D"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c", "d", "e", "f"},
			"A": []string{"c", "d", "e", "f", ""},
			"B": []string{"c", "d", "e", "f", ""},
			"C": []string{"c", "d", "e", "f", ""},
			"D": []string{"c", "d", "e", "f", ""},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "f"},
			"B": []string{"c"},
			"C": []string{"d"},
			"D": []string{"e"},
		},
	},
	{
		tgCycle1, true, 1, 2, 3,
		[]string{"S"},
		[]string{"a", "b"},
		[]string{"S"}, []string{"S"},
		[]string{"S"}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b"},
		},
		map[string][]string{
			"S": []string{"$"},
		},
	},
	{
		tgCycle2, true, 2, 4, 6,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"A"}, []string{"A"},
		[]string{"A"}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c", "d"},
			"A": []string{"c", "d"},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
		},
	},
	{
		tgCycle3, true, 2, 4, 6,
		[]string{"S", "A"},
		[]string{"a", "b", "c", "d"},
		[]string{"S", "A"}, []string{},
		[]string{"S", "A"}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c", "d"},
			"A": []string{"a", "b", "c", "d"},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
		},
	},
	{
		tgCycle4, true, 3, 6, 9,
		[]string{"S", "A", "B"},
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"S", "A", "B"}, []string{},
		[]string{"S", "A", "B"}, []string{}, []string{},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", "c", "d", "e", "f"},
			"A": []string{"a", "b", "c", "d", "e", "f"},
			"B": []string{"a", "b", "c", "d", "e", "f"},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"$"},
			"B": []string{"$"},
		},
	},
	{
		tgNullable1, false, 7, 2, 10,
		[]string{"S", "A", "B", "C", "D", "E", "F"},
		[]string{"a", "b"},
		[]string{}, []string{},
		[]string{}, []string{"C", "D"}, []string{"S", "C", "D"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", ""},
			"A": []string{"a"},
			"B": []string{"b"},
			"C": []string{"a", ""},
			"D": []string{"b", ""},
			"E": []string{"a"},
			"F": []string{"b"},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"b"},
			"B": []string{"$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"$"},
		},
	},
	{
		tgNullable2, false, 7, 2, 12,
		[]string{"S", "A", "B", "C", "D", "F", "E"},
		[]string{"a", "b"},
		[]string{}, []string{},
		[]string{}, []string{"B", "D"}, []string{"S", "B", "C", "D"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", ""},
			"A": []string{"b"},
			"B": []string{"a", "b", ""},
			"C": []string{"a", "b", ""},
			"D": []string{"b", ""},
			"E": []string{"a"},
			"F": []string{"b"},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "b", "$"},
			"B": []string{"b", "$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"a", "b", "$"},
		},
	},
	{
		tgNullable3, true, 8, 2, 16,
		[]string{"S", "A", "B", "G", "C", "D", "F", "E"},
		[]string{"a", "b"},
		[]string{"S", "G"}, []string{},
		[]string{"S", "G"}, []string{"B", "D"},
		[]string{"S", "B", "G", "C", "D"},
		[]string{},
		[]string{},
		map[string][]string{
			"S": []string{"a", "b", ""},
			"A": []string{"b"},
			"B": []string{"a", "b", ""},
			"C": []string{"a", "b", ""},
			"D": []string{"b", ""},
			"E": []string{"a"},
			"F": []string{"b"},
			"G": []string{"a", "b", ""},
		},
		map[string][]string{
			"S": []string{"$"},
			"A": []string{"a", "b", "$"},
			"B": []string{"b", "$"},
			"C": []string{"b", "$"},
			"D": []string{"$"},
			"E": []string{"b", "$"},
			"F": []string{"a", "b", "$"},
			"G": []string{"$"},
		},
	},
	{
		tgUnreachable1, true, 5, 5, 8,
		[]string{"E", "T", "F", "Digits", "U"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
		[]string{"E", "T"}, []string{"E", "T"},
		[]string{}, []string{}, []string{},
		[]string{"U"},
		[]string{},
		map[string][]string{
			"F":      []string{"\\(", "[[:digit:]]+"},
			"T":      []string{"\\(", "[[:digit:]]+"},
			"E":      []string{"\\(", "[[:digit:]]+"},
			"Digits": []string{"[[:digit:]]+"},
			"U":      []string{"\\(", "[[:digit:]]+"},
		},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\*", "\\)", "$"},
			"E":      []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
			"U":      []string{},
		},
	},
	{
		tgUnreachable2, true, 7, 8, 14,
		[]string{"E", "T", "F", "Digits", "U", "V", "W"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+", "a", "b", "c"},
		[]string{"E", "T", "W"}, []string{"E", "T", "W"},
		[]string{"W"}, []string{"W"}, []string{"W"},
		[]string{"U", "V", "W"},
		[]string{},
		map[string][]string{
			"F":      []string{"\\(", "[[:digit:]]+"},
			"T":      []string{"\\(", "[[:digit:]]+"},
			"E":      []string{"\\(", "[[:digit:]]+"},
			"Digits": []string{"[[:digit:]]+"},
			"U":      []string{"\\(", "[[:digit:]]+"},
			"V":      []string{"a", "b"},
			"W":      []string{"a", "b", "c", ""},
		},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\*", "\\)", "$"},
			"E":      []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
			"U":      []string{},
			"V":      []string{},
			"W":      []string{},
		},
	},
	{
		tgUnproductive1, true, 5, 6, 8,
		[]string{"E", "T", "F", "Digits", "U"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+", "u"},
		[]string{"E", "T", "U"}, []string{"E", "T", "U"},
		[]string{}, []string{}, []string{},
		[]string{"U"},
		[]string{"U"},
		map[string][]string{
			"F":      []string{"\\(", "[[:digit:]]+"},
			"T":      []string{"\\(", "[[:digit:]]+"},
			"E":      []string{"\\(", "[[:digit:]]+"},
			"Digits": []string{"[[:digit:]]+"},
			"U":      []string{},
		},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\*", "\\)", "$"},
			"E":      []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
			"U":      []string{"u"},
		},
	},
	{
		tgUnproductive2, true, 7, 7, 11,
		[]string{"E", "T", "F", "Digits", "W", "U", "V"},
		[]string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+", "u", "w"},
		[]string{"E", "T", "W", "U", "V"},
		[]string{"E", "T", "W", "U", "V"},
		[]string{"V"}, []string{}, []string{},
		[]string{"U", "V"},
		[]string{"W", "U", "V"},
		map[string][]string{
			"F":      []string{"\\(", "[[:digit:]]+"},
			"T":      []string{"\\(", "[[:digit:]]+"},
			"E":      []string{"\\(", "[[:digit:]]+"},
			"Digits": []string{"[[:digit:]]+"},
			"U":      []string{},
			"V":      []string{},
			"W":      []string{},
		},
		map[string][]string{
			"F":      []string{"\\+", "\\*", "\\)", "$"},
			"T":      []string{"\\+", "\\*", "\\)", "$"},
			"E":      []string{"\\+", "\\)", "$"},
			"Digits": []string{"\\*", "\\+", "\\)", "$"},
			"U":      []string{"u"},
			"V":      []string{},
			"W":      []string{"\\*", "\\+", "\\)", "w", "$"},
		},
	},
	{
		tgAdventure, false, 9, 23, 29,
		[]string{"S", "Command", "Direction", "Character",
			"OptionalWeapon", "Openable", "OptionalOpener", "Weapon",
			"Opener"},
		[]string{"go", "kill", "open", "north", "east", "south", "west",
			"up", "down", "dwarf", "snake", "unicorn", "rat", "with",
			"sword", "poison", "hands", "door", "trapdoor", "box",
			"mind", "key", "scissors"},
		[]string{},
		[]string{},
		[]string{},
		[]string{"OptionalWeapon", "OptionalOpener"},
		[]string{"OptionalWeapon", "OptionalOpener"},
		[]string{},
		[]string{},
		map[string][]string{
			"S":              []string{"go", "kill", "open"},
			"Command":        []string{"go", "kill", "open"},
			"Direction":      []string{"north", "east", "south", "west", "down", "up"},
			"Character":      []string{"dwarf", "snake", "unicorn", "rat"},
			"OptionalWeapon": []string{"with", ""},
			"Openable":       []string{"door", "trapdoor", "box", "mind"},
			"OptionalOpener": []string{"with", ""},
			"Weapon":         []string{"dwarf", "sword", "poison", "hands"},
			"Opener":         []string{"sword", "key", "scissors"},
		},
		map[string][]string{
			"S":              []string{"$"},
			"Command":        []string{"$"},
			"Direction":      []string{"$"},
			"Character":      []string{"$", "with"},
			"OptionalWeapon": []string{"$"},
			"Openable":       []string{"$", "with"},
			"OptionalOpener": []string{"$"},
			"Weapon":         []string{"$"},
			"Opener":         []string{"$"},
		},
	},
}
