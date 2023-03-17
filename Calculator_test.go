package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestCalculation(t *testing.T) {
	testTable := []struct {
		mathematicalOperation string
		expectedValue         string
		expectedErr           error
	}{
		{
			mathematicalOperation: "1 + 2",
			expectedValue:         "3",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "1 / 0",
			expectedValue:         "",
			expectedErr:           fmt.Errorf("на ноль делить нельзя"),
		},
		{
			mathematicalOperation: "10 * 10",
			expectedValue:         "100",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "4 / 5",
			expectedValue:         "0",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "0 - 10",
			expectedValue:         "-10",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "V + IV",
			expectedValue:         "IX",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "VII * VIII",
			expectedValue:         "LVI",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "VII / II",
			expectedValue:         "III",
			expectedErr:           nil,
		},
		{
			mathematicalOperation: "VII / 2",
			expectedValue:         "",
			expectedErr:           errors.New("некорректная строка"),
		},
		{
			mathematicalOperation: "a + b",
			expectedValue:         "",
			expectedErr:           errors.New("некорректная строка"),
		},
		{
			mathematicalOperation: "1 + 2 + 3",
			expectedValue:         "",
			expectedErr:           errors.New("строка не является математической операцией"),
		},
		{
			mathematicalOperation: "III - III",
			expectedValue:         "",
			expectedErr:           errors.New("в римской системе есть только положительные числа"),
		},
		{
			mathematicalOperation: "III & III",
			expectedValue:         "",
			expectedErr:           errors.New("строка не является математической операцией"),
		},
	}
	for _, testCase := range testTable {
		result, err := Calculation(testCase.mathematicalOperation)
		if err == nil {
			t.Logf("%s, result %s", testCase.mathematicalOperation, result)
		} else {
			t.Logf("%s, result %s", testCase.mathematicalOperation, err)
		}

		if result != testCase.expectedValue {
			if err == nil {
				t.Errorf("Incorrect result. Expect %s, got %s",
					testCase.expectedValue, result)
			} else {
				t.Errorf("Incorrect result. Expect %s, got %s",
					testCase.expectedErr, err)
			}
		}
	}
}
