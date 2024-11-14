package test

import (
	"testing"

	"github.com/mgechev/revive/internal/ifelse"
	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early_return", &rule.EarlyReturnRule{})
	testRule(t, "early_return_scope", &rule.EarlyReturnRule{}, &lint.RuleConfig{Arguments: []any{ifelse.PreserveScope}})
}