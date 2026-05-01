package mathutil

import "testing"

func TestTest(t *testing.T) {
	got := 1
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
