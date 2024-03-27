package utility

import (
	"testing"
)

/*
	Testler aşağıdaki şekilde çalıştırılabilir

	go test .

	Bu durumda aşağıdaki test sonucu üretilir

	--- FAIL: TestAdd (0.00s)
    unittestops_test.go:9: got '\x19', wanted '\n'
	FAIL
	FAIL    example/unittestops     0.002s
	FAIL

	Tüm testleri görmek için  (başarılı olanlar da dahil) aşağıdaki gibi çalıştırabiliriz.
	v = verbose anlamındadır
	go test -v
*/

func TestAddFail(t *testing.T) {
	actual := Add(13, 12)
	expected := 10
	if actual != expected {
		t.Errorf("actual is %q, wanted %q", actual, expected)
	}
}

func TestAddSuccess(t *testing.T) {
	actual := Add(13, 12)
	expected := 25
	if actual != expected {
		t.Errorf("Actual is %q, expected is %q", actual, expected)
	}
}

func TestSign(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 1},
		{-1, -1},
		{0, 0},
	}

	for _, test := range tests {
		if output := Sign(test.input); output != test.expected {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, received: {%d}", test.input, test.expected, output)
		}
	}
}
