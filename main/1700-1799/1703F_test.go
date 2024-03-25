// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1703/F
// https://codeforces.com/problemset/status/1703/problem/F
func Test_cf1703F(t *testing.T) {
	testCases := [][2]string{
		{
			`5
8
1 1 2 3 8 2 1 4
2
1 2
10
0 2 1 6 3 4 1 2 8 3
2
1 1000000000
3
0 1000000000 2`,
			`3
0
10
0
1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1703F)
}