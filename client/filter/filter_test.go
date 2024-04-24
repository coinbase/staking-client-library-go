package filter_test

import (
	"testing"

	"github.com/coinbase/staking-client-library-go/client/filter"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/suite"
)

// This test suite is used to test the underlying filter package, which powers the other filter types,
// such as ListStakes and ListRewards.
func TestFilterTestSuite(t *testing.T) {
	suite.Run(t, new(FilterTestSuite))
}

type FilterTestSuite struct {
	suite.Suite
}

func (s *FilterTestSuite) TestFilter_HappyPath() {
	f := WithA().Gt(10).String()
	assert.Equal(s.T(), "a > 10", f)

	// Filter to perform a simple AND between 2 conditions.
	f = WithA().Gte(10).And(WithA().Lt(20)).String()
	assert.Equal(s.T(), "(a >= 10 AND a < 20)", f)

	// Filter to perform a simple AND between 3 conditions with no special order enforcement.
	f = WithA().Gte(10).And(WithA().Lt(20)).And(WithB().Eq("example")).String()
	assert.Equal(s.T(), "(a >= 10 AND a < 20 AND b = 'example')", f)

	// Filter to perform a simple AND between 3 conditions while enforcing order of operations - forcing latter 2 conditions to be AND'ed first.
	f = WithA().Gte(10).And(WithA().Lt(20).And(WithB().Eq("example"))).String()
	assert.Equal(s.T(), "(a >= 10 AND (a < 20 AND b = 'example'))", f)
}

// AFilter is a custom type to define filter operation on the 'a' field.
type AFilter struct{}

// Eq method is a custom method to define the equals operation on the 'a' field.
func (AFilter) Eq(value int) *filter.Term {
	return filter.NewTerm("a", filter.Equals, value)
}

// Neq method is a custom method to define the not equals operation on the 'a' field.
func (AFilter) Neq(value int) *filter.Term {
	return filter.NewTerm("a", filter.NotEquals, value)
}

// Gt method is a custom method to define the greater than operation on the 'a' field.
func (AFilter) Gt(value int) *filter.Term {
	return filter.NewTerm("a", filter.GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the 'a' field.
func (AFilter) Gte(value int) *filter.Term {
	return filter.NewTerm("a", filter.GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the 'a' field.
func (AFilter) Lt(value int) *filter.Term {
	return filter.NewTerm("a", filter.LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the 'a' field.
func (AFilter) Lte(value int) *filter.Term {
	return filter.NewTerm("a", filter.LessEquals, value)
}

func WithA() AFilter {
	return AFilter{}
}

// BFilter is a custom type to define filter operation on the 'b' field.
type BFilter struct{}

// Eq method is a custom method to define the equals operation on the 'b' field.
func (BFilter) Eq(value string) *filter.Term {
	return filter.NewTerm("b", filter.Equals, value)
}

func WithB() BFilter {
	return BFilter{}
}
