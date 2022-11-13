package path

import (
	"testing"
)

func TestPath(t *testing.T) {
	p := Wrap("/parent/stem.suffix")
	if p.Full != "/parent/stem.suffix" {
		t.Fail()
	}
	if p.Name != "stem.suffix" {
		t.Fail()
	}
	if p.Parent.Full != "/parent" {
		t.Fail()
	}
	if p.Stem != "stem" {
		t.Fail()
	}
	if p.Suffix != ".suffix" {
		t.Fail()
	}
}
