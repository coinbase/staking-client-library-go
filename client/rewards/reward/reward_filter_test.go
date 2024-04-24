package reward_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/coinbase/staking-client-library-go/client/rewards/reward"
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
	f := reward.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithAddress() {
	f := reward.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithEpoch() {
	var f string

	f = reward.WithEpoch().Eq(10).String()
	assert.Equal(s.T(), "epoch = 10", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithEpoch().Lt(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch < 10)", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithEpoch().Lte(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch <= 10)", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithEpoch().Gt(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch > 10)", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithEpoch().Gte(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch >= 10)", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithEpoch().Neq(10)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND epoch != 10)", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithDate() {
	var f string

	f = reward.WithDate().Eq("2020-01-01").String()
	assert.Equal(s.T(), "date = '2020-01-01'", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithDate().Lt("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date < '2020-01-01')", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithDate().Lte("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date <= '2020-01-01')", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithDate().Gt("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date > '2020-01-01')", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithDate().Gte("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date >= '2020-01-01')", f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithDate().Neq("2020-01-01")).String()
	assert.Equal(s.T(), "(address = 'abcd' AND date != '2020-01-01')", f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_WithPeriodEndTime() {
	var f string

	now := time.Now()
	fiveDaysAgo := now.AddDate(0, 0, -5)

	nowRFC3339 := now.Format(time.RFC3339)
	fiveDaysAgoRFC3339 := fiveDaysAgo.Format(time.RFC3339)

	f = reward.WithPeriodEndTime().Eq(fiveDaysAgo).String()
	assert.Equal(s.T(), fmt.Sprintf("period_end_time = '%s'", fiveDaysAgoRFC3339), f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithPeriodEndTime().Lt(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time < '%s')", fiveDaysAgoRFC3339), f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithPeriodEndTime().Lte(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time <= '%s')", fiveDaysAgoRFC3339), f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithPeriodEndTime().Gt(fiveDaysAgo)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time > '%s')", fiveDaysAgoRFC3339), f)

	f = reward.WithAddress().Eq("abcd").And(reward.WithPeriodEndTime().Gte(fiveDaysAgo)).And(reward.WithPeriodEndTime().Lt(now)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time >= '%s' AND period_end_time < '%s')", fiveDaysAgoRFC3339, nowRFC3339), f)
}

func (s *ListRewardsFilterSuite) TestListStakesFilter_MultipleDifferentComparisons() {
	var f string

	now := time.Now()
	fiveDaysAgo := now.AddDate(0, 0, -5)

	nowRFC3339 := now.Format(time.RFC3339)
	fiveDaysAgoRFC3339 := fiveDaysAgo.Format(time.RFC3339)

	f = reward.WithAddress().Eq("abcd").And(reward.WithPeriodEndTime().Gte(fiveDaysAgo)).And(reward.WithPeriodEndTime().Lt(now)).And(reward.WithEpoch().Eq(50)).String()
	assert.Equal(s.T(), fmt.Sprintf("(address = 'abcd' AND period_end_time >= '%s' AND period_end_time < '%s' AND epoch = 50)", fiveDaysAgoRFC3339, nowRFC3339), f)
}
