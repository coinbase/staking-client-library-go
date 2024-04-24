package rewards_filter_test

import (
	"fmt"
	"testing"
	"time"

	filter "github.com/coinbase/staking-client-library-go/client/rewards/rewards_filter"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/suite"
)

func TestListRewardsFilterSuite(t *testing.T) {
	suite.Run(t, new(ListRewardsFilterSuite))
}

type ListRewardsFilterSuite struct {
	suite.Suite
}

func (s *ListRewardsFilterSuite) TestListRewardsFilter_WithAddress() {
	f := filter.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithAddress() {
	f := filter.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithEpoch() {
	var f string

	f = filter.WithEpoch().Eq(10).String()
	assert.Equal(s.T(), "epoch = 10", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEpoch().Lt(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch < 10)", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEpoch().Lte(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch <= 10)", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEpoch().Gt(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch > 10)", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEpoch().Gte(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch >= 10)", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEpoch().Neq(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch != 10)", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithDate() {
	var f string

	f = filter.WithDate().Eq("2020-01-01").String()
	assert.Equal(s.T(), "date = '2020-01-01'", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithDate().Lt("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date < '2020-01-01')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithDate().Lte("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date <= '2020-01-01')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithDate().Gt("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date > '2020-01-01')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithDate().Gte("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date >= '2020-01-01')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithDate().Neq("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date != '2020-01-01')", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithPeriodEndTime() {
	var f string

	now := time.Now()
	fiveDaysAgo := now.AddDate(0, 0, -5)

	nowRFC3339 := now.Format(time.RFC3339)
	fiveDaysAgoRFC3339 := fiveDaysAgo.Format(time.RFC3339)

	f = filter.WithPeriodEndTime().Eq(fiveDaysAgo).String()
	assert.Equal(s.T(), fmt.Sprintf("period_end_time = '%s'", fiveDaysAgoRFC3339), f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Lt(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time < '%s')", fiveDaysAgoRFC3339), f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Lte(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time <= '%s')", fiveDaysAgoRFC3339), f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Gt(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time > '%s')", fiveDaysAgoRFC3339), f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Gte(fiveDaysAgo)).And(filter.WithPeriodEndTime().Lt(now)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time >= '%s' AND period_end_time < '%s')", fiveDaysAgoRFC3339, nowRFC3339), f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Gte(fiveDaysAgo).And(filter.WithPeriodEndTime().Lt(now))).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND (period_end_time >= '%s' AND period_end_time < '%s'))", fiveDaysAgoRFC3339, nowRFC3339), f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_MultipleDifferentComparisons() {
	var f string

	now := time.Now()
	fiveDaysAgo := now.AddDate(0, 0, -5)

	nowRFC3339 := now.Format(time.RFC3339)
	fiveDaysAgoRFC3339 := fiveDaysAgo.Format(time.RFC3339)

	f = filter.WithAddress().Eq("abcd").And(filter.WithPeriodEndTime().Gte(fiveDaysAgo)).And(filter.WithPeriodEndTime().Lt(now)).And(filter.WithEpoch().Eq(50)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time >= '%s' AND period_end_time < '%s' AND epoch = 50)", fiveDaysAgoRFC3339, nowRFC3339), f)
}
