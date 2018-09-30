package slrp_test

type testCase struct {
	filename string
	accepts  []string
	rejects  []string
}

var parserTestCases = []testCase{
	{
		tgArithLr,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgArithNlr,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgBalParens2,
		[]string{"", "()", "(())", "(()())"},
		[]string{"(", ")", "(()"},
	},
	{
		tgZeroOne,
		[]string{"01", "0011"},
		[]string{"", "0", "1", "001", "001011"},
	},
	{
		tgIndirectLr1,
		[]string{"a", "b", "ada", "ca", "ccca", "bda"},
		[]string{"cb", "da", "ad", "bd"},
	},
	{
		tgIndirectLr2,
		[]string{"a", "b", "ca", "dca", "aedcaedca"},
		[]string{"cb", "da", "ad", "bd"},
	},
	{
		tgIndirectLr3,
		[]string{"a", "b", "ca", "dca", "fedcfedca"},
		[]string{"cb", "da", "ad", "bd"},
	},
	{
		tgUnreachable1,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgUnreachable2,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgUnproductive1,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
	},
	{
		tgUnproductive2,
		[]string{"3", "3+4", "3+4*5", "(3+4)*5", "  (3 +   4) * 5  "},
		[]string{"", "+", "3+", "-3", "(3+4*5"},
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
	{
		tgSlrE,
		[]string{"ab"},
		[]string{"a", "b", ""},
	},
}
