package datastruct_test

import (
	"github.com/paulgriffiths/contextfree/datastruct"
	"testing"
)

var (
	siE    = datastruct.NewSetInt()
	si0    = datastruct.NewSetInt(0)
	si1    = datastruct.NewSetInt(1)
	si2    = datastruct.NewSetInt(2)
	si4    = datastruct.NewSetInt(4)
	si01   = datastruct.NewSetInt(0, 1)
	si10   = datastruct.NewSetInt(1, 0)
	si02   = datastruct.NewSetInt(0, 2)
	si12   = datastruct.NewSetInt(1, 2)
	si23   = datastruct.NewSetInt(2, 3)
	si34   = datastruct.NewSetInt(3, 4)
	si012  = datastruct.NewSetInt(0, 1, 2)
	si123  = datastruct.NewSetInt(1, 2, 3)
	si1234 = datastruct.NewSetInt(1, 2, 3, 4)
)

func TestSetIntEquals(t *testing.T) {
	testCases := []struct {
		a, b  datastruct.SetInt
		equal bool
	}{
		{siE, siE, true}, {siE, si0, false}, {si0, siE, false},
		{si0, si0, true}, {si0, si1, false}, {si1, si0, false},
		{si1, si1, true}, {si01, si10, true}, {si01, si02, false},
		{si012, si01, false},
	}

	for n, tc := range testCases {
		if r := tc.a.Equals(tc.b); r != tc.equal {
			t.Errorf("case %d, got %t, want %t", n+1, r, tc.equal)
		}
	}
}

func TestSetIntLength(t *testing.T) {
	testCases := []struct {
		values []int
		length int
	}{
		{[]int{1}, 1}, {[]int{1, 2}, 2}, {[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 3}, {[]int{1, 3, 3}, 2}, {[]int{2, 2, 2}, 1},
	}

	for n, tc := range testCases {
		s := datastruct.NewSetInt(tc.values...)
		if r := s.Length(); r != tc.length {
			t.Errorf("case %d, got %d, want %d", n+1, r, tc.length)
		}
	}
}

func TestSetIntContains(t *testing.T) {
	testCases := []struct {
		values   []int
		contains []bool
	}{
		{[]int{0}, []bool{true, false, false, false, false}},
		{[]int{1}, []bool{false, true, false, false, false}},
		{[]int{2}, []bool{false, false, true, false, false}},
		{[]int{3}, []bool{false, false, false, true, false}},
		{[]int{4}, []bool{false, false, false, false, true}},
		{[]int{0, 1}, []bool{true, true, false, false, false}},
		{[]int{0, 0, 1}, []bool{true, true, false, false, false}},
		{[]int{0, 0, 0, 1}, []bool{true, true, false, false, false}},
		{[]int{1, 2, 1}, []bool{false, true, true, false, false}},
		{[]int{3, 3, 2}, []bool{false, false, true, true, false}},
		{[]int{4, 3, 2}, []bool{false, false, true, true, true}},
		{[]int{4, 2, 2}, []bool{false, false, true, false, true}},
	}

	for n, tc := range testCases {
		s := datastruct.NewSetInt(tc.values...)
		for m, c := range tc.contains {
			if r := s.Contains(m); r != c {
				t.Errorf("case (%d,%d), got %t, want %t", n+1, m+1, r, c)
			}
		}
	}
}

func TestSetIntUnion(t *testing.T) {
	testCases := []struct {
		a, b, u datastruct.SetInt
	}{
		{si12, si34, si1234}, {si12, si23, si123}, {si12, si12, si12},
		{si12, si1, si12}, {si12, siE, si12}, {siE, siE, siE},
	}

	for n, tc := range testCases {
		if s := tc.a.Union(tc.b); !s.Equals(tc.u) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.u)
		}
	}
}

func TestSetIntIntersection(t *testing.T) {
	testCases := []struct {
		a, b, i datastruct.SetInt
	}{
		{si12, si34, siE}, {si12, si23, si2}, {si12, si12, si12},
		{si12, si1, si1}, {si12, siE, siE}, {siE, siE, siE},
	}

	for n, tc := range testCases {
		if s := tc.a.Intersection(tc.b); !s.Equals(tc.i) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.i)
		}
	}
}

func TestSetIntDifference(t *testing.T) {
	testCases := []struct {
		a, b, i datastruct.SetInt
	}{
		{si1234, si34, si12}, {si1234, si123, si4},
		{si012, si12, si0}, {si1234, si1234, siE},
	}

	for n, tc := range testCases {
		if s := tc.a.Difference(tc.b); !s.Equals(tc.i) {
			t.Errorf("case %d, got %v, want %v", n+1, s, tc.i)
		}
	}
}
