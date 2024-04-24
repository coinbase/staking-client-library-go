package stakes_filter

import (
	"time"

	"github.com/coinbase/staking-client-library-go/client/filter"
)

// AddressFilter is a custom type to define filter operation on the address field.
type AddressFilter struct{}

// Eq method is a custom method to define the equals operation on the address field.
func (AddressFilter) Eq(value string) *filter.Term {
	return filter.NewTerm("address", filter.Equals, value)
}

// WithAddress creates a new filter block that has `address` as the key.
// Needs a companion operator to be functional (ex: WithAddress().Eq("my_address").
func WithAddress() AddressFilter {
	return AddressFilter{}
}

// EvaluationTimeFilter is a custom type to define filter operation on the evaluation_time field.
type EvaluationTimeFilter struct{}

// Eq method is a custom method to define the equals operation on the evaluation_time field.
func (EvaluationTimeFilter) Eq(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.Equals, value)
}

// Neq method is a custom method to define the not equals operation on the evaluation_time field.
func (EvaluationTimeFilter) Neq(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.NotEquals, value)
}

// Gt method is a custom method to define the greater than operation on the evaluation_time field.
func (EvaluationTimeFilter) Gt(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the evaluation_time field.
func (EvaluationTimeFilter) Gte(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the evaluation_time field.
func (EvaluationTimeFilter) Lt(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the evaluation_time field.
func (EvaluationTimeFilter) Lte(value time.Time) *filter.Term {
	return filter.NewTerm("evaluation_time", filter.LessEquals, value)
}

// Instructs the backend API to return rewards aggregations that have `evaluation_time` set in a manner which matches the desired filter.
// Needs a companion comparison operator (ex: >, <, =, etc) to be functional.
func WithEvaluationTime() EvaluationTimeFilter {
	return EvaluationTimeFilter{}
}
