package unittestops

import "testing"

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

func TestAdd(t *testing.T) {
	got := Add(13, 12)
	want := 10
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAdd2(t *testing.T) {
	got := Add(13, 12)
	want := 25
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
