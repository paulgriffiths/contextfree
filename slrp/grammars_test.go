package slrp

var tgFilePrefix = "../testdata/grammars/"

var tgArithLr = tgFilePrefix + "01_arith_lr.grammar"
var tgArithNlr = tgFilePrefix + "02_arith_nlr.grammar"
var tgArithAmbig = tgFilePrefix + "03_arith_ambig.grammar"
var tgBalParens1 = tgFilePrefix + "04_bal_parens_1.grammar"
var tgBalParens2 = tgFilePrefix + "05_bal_parens_2.grammar"
var tgZeroOne = tgFilePrefix + "06_zero_one.grammar"
var tgIndirectLr1 = tgFilePrefix + "07_indirect_lr_1.grammar"
var tgIndirectLr2 = tgFilePrefix + "08_indirect_lr_2.grammar"
var tgIndirectLr3 = tgFilePrefix + "09_indirect_lr_3.grammar"
var tgCycle1 = tgFilePrefix + "10_cycle_1.grammar"
var tgCycle2 = tgFilePrefix + "11_cycle_2.grammar"
var tgCycle3 = tgFilePrefix + "12_cycle_3.grammar"
var tgCycle4 = tgFilePrefix + "13_cycle_4.grammar"
var tgNullable1 = tgFilePrefix + "14_nullable_1.grammar"
var tgNullable2 = tgFilePrefix + "15_nullable_2.grammar"
var tgNullable3 = tgFilePrefix + "16_nullable_3.grammar"
var tgUnreachable1 = tgFilePrefix + "17_unreachable_1.grammar"
var tgUnreachable2 = tgFilePrefix + "18_unreachable_2.grammar"
var tgUnproductive1 = tgFilePrefix + "19_unproductive_1.grammar"
var tgUnproductive2 = tgFilePrefix + "20_unproductive_2.grammar"
var tgAdventure = tgFilePrefix + "21_adventure.grammar"

var tgOutArithLrRaw = tgFilePrefix + "output/01_arith_lr_raw.grammar"
var tgOutArithNlrRaw = tgFilePrefix + "output/02_arith_nlr_raw.grammar"
var tgOutArithAmbigRaw = tgFilePrefix + "output/03_arith_ambig_raw.grammar"
var tgOutBalParens1Raw = tgFilePrefix + "output/04_bal_parens_1_raw.grammar"
var tgOutBalParens2Raw = tgFilePrefix + "output/05_bal_parens_2_raw.grammar"
var tgOutZeroOneRaw = tgFilePrefix + "output/06_zero_one_raw.grammar"

var tgBadMissingHead1 = tgFilePrefix + "bad/missing_head_1.grammar"
var tgBadMissingBody1 = tgFilePrefix + "bad/missing_body_1.grammar"
var tgBadMissingBody2 = tgFilePrefix + "bad/missing_body_2.grammar"
var tgBadMissingBody3 = tgFilePrefix + "bad/missing_body_3.grammar"
var tgBadMissingBody4 = tgFilePrefix + "bad/missing_body_4.grammar"
var tgBadENotAlone1 = tgFilePrefix + "bad/e_not_alone_1.grammar"
var tgBadENotAlone2 = tgFilePrefix + "bad/e_not_alone_2.grammar"
var tgBadMissingArrow1 = tgFilePrefix + "bad/missing_arrow_1.grammar"

var tgBadUnterminatedTerminal1 = tgFilePrefix + "bad/unterminated_terminal_1.grammar"
var tgBadUnterminatedTerminal2 = tgFilePrefix + "bad/unterminated_terminal_2.grammar"
var tgBadIllegalCharacter1 = tgFilePrefix + "bad/illegal_character_1.grammar"
var tgBadIllegalCharacter2 = tgFilePrefix + "bad/illegal_character_2.grammar"

var testGrammars = []string{
	tgArithLr,
	tgArithNlr,
	tgArithAmbig,
	tgBalParens1,
	tgBalParens2,
	tgZeroOne,
	tgIndirectLr1,
	tgIndirectLr2,
	tgIndirectLr3,
	tgCycle1,
	tgCycle2,
	tgCycle3,
	tgCycle4,
	tgNullable1,
	tgNullable2,
	tgNullable3,
	tgUnreachable1,
	tgUnreachable2,
	tgUnproductive1,
	tgUnproductive2,
	tgAdventure,
}
