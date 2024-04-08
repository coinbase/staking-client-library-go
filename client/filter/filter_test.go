package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AFilter is a custom type to define filter operation on the 'a' field.
type AFilter struct{}

// Eq method is a custom method to define the equals operation on the 'a' field.
func (AFilter) Eq(value int) *Term {
	return NewTerm("a", Equals, value)
}

// Neq method is a custom method to define the not equals operation on the 'a' field.
func (AFilter) Neq(value int) *Term {
	return NewTerm("a", NotEquals, value)
}

// Gt method is a custom method to define the greater than operation on the 'a' field.
func (AFilter) Gt(value int) *Term {
	return NewTerm("a", GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the 'a' field.
func (AFilter) Gte(value int) *Term {
	return NewTerm("a", GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the 'a' field.
func (AFilter) Lt(value int) *Term {
	return NewTerm("a", LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the 'a' field.
func (AFilter) Lte(value int) *Term {
	return NewTerm("a", LessEquals, value)
}

func WithA() AFilter {
	return AFilter{}
}

// BFilter is a custom type to define filter operation on the 'b' field.
type BFilter struct{}

// Eq method is a custom method to define the equals operation on the 'b' field.
func (BFilter) Eq(value string) *Term {
	return NewTerm("b", Equals, value)
}

func WithB() BFilter {
	return BFilter{}
}

func TestFilter(t *testing.T) {
	f := WithA().Gt(10).String()
	assert.Equal(t, "a > '10'", f)

	// Filter to perform a simple AND between 2 conditions.
	f = WithA().Gte(10).And(WithA().Lt(20)).String()
	assert.Equal(t, "(a >= '10' AND a < '20')", f)

	// Filter to perform a simple AND between 3 conditions with no special order enforcement.
	f = WithA().Gte(10).And(WithA().Lt(20)).And(WithB().Eq("example")).String()
	assert.Equal(t, "(a >= '10' AND a < '20' AND b = 'example')", f)

	// Filter to perform a simple AND between 3 conditions while enforcing order of operations - forcing latter 2 conditions to be AND'ed first.
	f = WithA().Gte(10).And(WithA().Lt(20).And(WithB().Eq("example"))).String()
	assert.Equal(t, "(a >= '10' AND (a < '20' AND b = 'example'))", f)
}
