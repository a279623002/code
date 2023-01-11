package strategy

import "testing"

func TestNewLogManager(t *testing.T) {
	f := &FileLogger{}
	lm := NewLogManager(f)
	f.Error()
	f.Info()
	d := &DbLogger{}
	lm = NewLogManager(d)
	lm.Error()
	lm.Info()
}
