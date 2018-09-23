package pp_test

type testCase struct {
	filename string
	accepts  []string
	rejects  []string
}

var parserTestCases = []testCase{
	{
		tgArithNlr,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgBalParens2,
		[]string{"", "()", "(())", "((()))", "(()())", "(()()())((()))"},
		[]string{"(", ")", "(()", "())", ")(", "(((((()))))()"},
	},
	{
		tgZeroOne,
		[]string{"01", "0011", "000111", "0 0 01 1 1"},
		[]string{"", "0", "1", "001", "011", "10", "111000", "001011",
			"0000001111101", "0 0 0 1 1 1"},
	},
	{
		tgAdventure,
		[]string{
			"go north",
			"go up",
			"kill dwarf",
			"kill rat with poison",
			"kill snake with hands",
			"kill unicorn with dwarf", // Deliberate
			"open door",
			"open door with key",
			"open trapdoor with sword",
			"open box with scissors",
			"open mind with scissors",
		},
		[]string{
			"go nowhere",
			"go crazy",
			"kill door",
			"kill snake with unicorn",
			"kill east with key",
			"open door with dwarf",
			"open scissors with key",
			"open trapdoor with poison",
			"sing about gold with thorin",
		},
	},
}
