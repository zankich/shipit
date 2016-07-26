package shipit

import "testing"

func TestAdd(t *testing.T) {
	want := 4
	got := Add(2, 2)
	if got != want {
		t.Fatal("want:", want, "got:", got)
	}
}
