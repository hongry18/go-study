package main

/**
 * filename_test.go
 * 파일명 suffix 로 _test를 붙여줘야 단위테스트가 가능하다
 * 그리고 파일 내부 함수로는 TestXxxx 형태의 PascalCase를 사용한다
 * arguments 는 *testing.T 를 선언한다
 */

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("god %d, want %d", ans, tt.want)
			}
		})
	}
}
