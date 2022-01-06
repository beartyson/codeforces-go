// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[2,2,1,2]`, 
			`3`,
		},
		{
			`[1,2,0]`, 
			`3`,
		},
		{
			`[3,0,1,4,1]`, 
			`4`,
		},
		{
			`[1,2,3,4,5,6,3,8,9,10,11,8]`,
			`4`,
		},
		{
			`[1,0,0,2,1,4,7,8,9,6,7,10,8]`,
			`6`,
		},
		{
			`[1,0,3,2,5,6,7,4,9,8,11,10,11,12,10]`,
			`11`,
		},
	}
	targetCaseNum :=  0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maximumInvitations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-274/problems/maximum-employees-to-be-invited-to-a-meeting/