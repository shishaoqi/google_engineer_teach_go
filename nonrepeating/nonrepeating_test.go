package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	} {
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑灰化肥灰会挥发发灰黑讳为黑灰花会飞", 7},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)

		if actual != tt.ans {
			t.Error("got %d for input %s; " + "expect %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑灰化肥灰会挥发发灰黑讳为黑灰花会飞"
	ans := 7

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; " +
				"expected %d", actual, s, ans)
		}
	}
}
