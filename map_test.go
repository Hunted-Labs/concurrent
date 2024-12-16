package concurrent_test

import (
	"testing"

	"github.com/Hunted-Labs/concurrent"
)

func TestMap_Load(t *testing.T) {
	m := concurrent.NewMap()
	m.Store("hello", "world")
	value, found := m.Load("hello")
	if !found {
		t.Fail()
	}
	if value != "world" {
		t.Fail()
	}
}
