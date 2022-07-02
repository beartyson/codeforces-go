// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`6
1 2 2 1 5
2 1 1 1 1 1`,
			`3`,
		},
		{
			`7
1 1 2 3 1 4
3 3 1 1 1 2 3`,
			`5`,
		},
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
