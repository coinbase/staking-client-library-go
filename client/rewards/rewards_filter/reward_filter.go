package rewards_filter

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

// WithAddress instructs the backend API to return rewards aggregations that have `address` set in a manner which matches the desired filter.
// Needs a companion equals operator to be functional (ex: WithAddress().Eq("my_address").
func WithAddress() AddressFilter {
	return AddressFilter{}
}

// EpochFilter is a custom type to define filter operation on the epoch field.
type EpochFilter struct{}

// Eq method is a custom method to define the equals operation on the epoch field.
func (EpochFilter) Eq(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.Equals, value)
}

// Neq method is a custom method to define the not equals operation on the epoch field.
func (EpochFilter) Neq(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.NotEquals, value)
}

// Gt method is a custom method to define the greater than operation on the epoch field.
func (EpochFilter) Gt(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the epoch field.
func (EpochFilter) Gte(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the epoch field.
func (EpochFilter) Lt(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the epoch field.
func (EpochFilter) Lte(value int) *filter.Term {
	return filter.NewTerm("epoch", filter.LessEquals, value)
}

// WithEpoch instructs the backend API to return rewards aggregations that have `epoch` set in a manner which matches the desired filter.
// Needs a companion comparison operator (ex: >, <, =, etc) to be functional.
func WithEpoch() EpochFilter {
	return EpochFilter{}
}

// DateFilter is a custom type to define filter operation on the date.
type DateFilter struct{}

// Eq method is a custom method to define the equals operation on the date.
func (DateFilter) Eq(value string) *filter.Term {
	return filter.NewTerm("date", filter.Equals, value)
}

// Neq method is a custom method to define the not equals operation on the date.
func (DateFilter) Neq(value string) *filter.Term {
	return filter.NewTerm("date", filter.NotEquals, value)
}

// Gt method is a custom method to define the greater than operation on the date.
func (DateFilter) Gt(value string) *filter.Term {
	return filter.NewTerm("date", filter.GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the date.
func (DateFilter) Gte(value string) *filter.Term {
	return filter.NewTerm("date", filter.GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the date.
func (DateFilter) Lt(value string) *filter.Term {
	return filter.NewTerm("date", filter.LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the date.
func (DateFilter) Lte(value string) *filter.Term {
	return filter.NewTerm("date", filter.LessEquals, value)
}

// WithDate instructs the backend API to return rewards aggregations that have `date` set in a manner which matches the desired filter.
// Needs a companion comparison operator (ex: >, <, =, etc) to be functional.
func WithDate() DateFilter {
	return DateFilter{}
}

// PeriodEndTimeFilter is a custom type to define filter operation on the period_end_time field.
type PeriodEndTimeFilter struct{}

// Eq method is a custom method to define the equals operation on the period_end_time field.
func (PeriodEndTimeFilter) Eq(value time.Time) *filter.Term {
	return filter.NewTerm("period_end_time", filter.Equals, value)
}

// Gt method is a custom method to define the greater than operation on the period_end_time field.
func (PeriodEndTimeFilter) Gt(value time.Time) *filter.Term {
	return filter.NewTerm("period_end_time", filter.GreaterThan, value)
}

// Gte method is a custom method to define the greater than or equals operation on the period_end_time field.
func (PeriodEndTimeFilter) Gte(value time.Time) *filter.Term {
	return filter.NewTerm("period_end_time", filter.GreaterEquals, value)
}

// Lt method is a custom method to define the less than operation on the period_end_time field.
func (PeriodEndTimeFilter) Lt(value time.Time) *filter.Term {
	return filter.NewTerm("period_end_time", filter.LessThan, value)
}

// Lte method is a custom method to define the less than or equals operation on the period_end_time field.
func (PeriodEndTimeFilter) Lte(value time.Time) *filter.Term {
	return filter.NewTerm("period_end_time", filter.LessEquals, value)
}

// Instructs the backend API to return rewards aggregations that have `period_end_time` set in a manner which matches the desired filter.
// Needs a companion comparison operator (ex: >, <, =, etc) to be functional.
func WithPeriodEndTime() PeriodEndTimeFilter {
	return PeriodEndTimeFilter{}
}
