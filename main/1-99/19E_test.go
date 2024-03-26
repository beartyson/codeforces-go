// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/19/E
// https://codeforces.com/problemset/status/19/problem/E
func Test_cf19E(t *testing.T) {
	testCases := [][2]string{
		{
			`4 4
1 2
1 3
2 4
3 4`,
			`4
1 2 3 4`,
		},
		{
			`4 5
1 2
2 3
3 4
4 1
1 3`,
			`1
5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf19E)
}
