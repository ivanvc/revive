package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

// String-of-int rule.
func TestStringOfInt(t *testing.T) {
	testRule(t, "string_of_int", &rule.StringOfIntRule{})
}