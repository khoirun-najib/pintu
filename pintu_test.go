package pintu

import "testing"

func TestMaxProfit(t *testing.T) {
	expectation := 299595961
	actual := MaxProfit()
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestLexicograph(t *testing.T) {
	expectation := "sfgxclryzidpuvejaqbtwmhkno"
	actual := Lexicograph()
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestSmallestOrder(t *testing.T) {
	expectation := "apqflbgtdcwijemyzrvuhksnxo"
	actual := SmallestLexOrder()
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFactorial(t *testing.T) {
	expectation := 19
	actual := RangeFactor(128)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func BenchmarkFactorial(b *testing.B) {
	input := 262144
	for i := 0; i < b.N; i++ {
		RangeFactor(input)
	}
}
