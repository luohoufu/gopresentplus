package pokemon

import (
	"testing"
)

func TestWeakAgainst(t *testing.T) {

	tests := []struct {
		First    Type
		Second   Type
		Expected bool
	}{
		{Grass, Fire, true},
		{Grass, Bug, true},
		{Fire, Water, true},
	}

	t.Log("Testing types")
	for _, test := range tests {
		if r := test.First.WeakAgainst(test.Second); r != test.Expected {
			t.Failf("Test %s against %s. Expected %t, got %t\n", test.First, test.Second, test.Expected, r)
		}
	}
}
