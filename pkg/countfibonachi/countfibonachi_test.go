package countfibonachi

import "testing"

func TestCountFibonachi(t *testing.T) {
	got := CountFibonachi(6)
	want := 8
	if got != want {
		t.Error("Expect ", want, "  got %v ", got)

	}
}
