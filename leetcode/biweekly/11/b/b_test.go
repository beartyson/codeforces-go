// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[[10,50],[60,120],[140,210]]`, `[[0,15],[60,70]]`, `8`, 
			`[60,68]`,
		},
		{
			`[[10,50],[60,120],[140,210]]`, `[[0,15],[60,70]]`, `12`, 
			`[]`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minAvailableDuration, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-11/problems/meeting-scheduler/
