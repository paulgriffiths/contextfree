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
		tgSlrE,
		[]string{"ab"},
		[]string{"a", "b", ""},
	},
}
