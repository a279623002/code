package decorator

import "testing"

func TestRun(t *testing.T) {
	if desc, cost := Run(); desc != "cmw" || cost != 3 {
		t.Errorf("want cmw and 3, but got %s %d", desc, cost)
	}
}
