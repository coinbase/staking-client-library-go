package v1

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRewardFilter(t *testing.T) {
	f := WithAddress().Eq("abcd").String()
	assert.Equal(t, "address = 'abcd'", f)

	f = WithEpoch().Eq(10).String()
	assert.Equal(t, "epoch = 10", f)

	f = WithAddress().Eq("abcd").And(WithEpoch().Lt(10)).String()
	assert.Equal(t, "(address = 'abcd' AND epoch < 10)", f)

	f = WithAddress().Eq("abcd").And(WithEpoch().Lte(10)).String()
	assert.Equal(t, "(address = 'abcd' AND epoch <= 10)", f)

	f = WithAddress().Eq("abcd").And(WithEpoch().Gt(10)).String()
	assert.Equal(t, "(address = 'abcd' AND epoch > 10)", f)

	f = WithAddress().Eq("abcd").And(WithEpoch().Gte(10)).String()
	assert.Equal(t, "(address = 'abcd' AND epoch >= 10)", f)

	f = WithAddress().Eq("abcd").And(WithEpoch().Neq(10)).String()
	assert.Equal(t, "(address = 'abcd' AND epoch != 10)", f)

	f = WithDate().Eq("2020-01-01").String()
	assert.Equal(t, "date = '2020-01-01'", f)

	f = WithAddress().Eq("abcd").And(WithDate().Lt("2020-01-01")).String()
	assert.Equal(t, "(address = 'abcd' AND date < '2020-01-01')", f)

	f = WithAddress().Eq("abcd").And(WithDate().Lte("2020-01-01")).String()
	assert.Equal(t, "(address = 'abcd' AND date <= '2020-01-01')", f)

	f = WithAddress().Eq("abcd").And(WithDate().Gt("2020-01-01")).String()
	assert.Equal(t, "(address = 'abcd' AND date > '2020-01-01')", f)

	f = WithAddress().Eq("abcd").And(WithDate().Gte("2020-01-01")).String()
	assert.Equal(t, "(address = 'abcd' AND date >= '2020-01-01')", f)

	f = WithAddress().Eq("abcd").And(WithDate().Neq("2020-01-01")).String()
	assert.Equal(t, "(address = 'abcd' AND date != '2020-01-01')", f)

	now := time.Now()
	fiveDaysAgo := now.AddDate(0, 0, -5)

	nowRFC3339 := now.Format(time.RFC3339)
	fiveDaysAgoRFC3339 := fiveDaysAgo.Format(time.RFC3339)

	f = WithPeriodEndTime().Eq(fiveDaysAgo).String()
	assert.Equal(t, fmt.Sprintf("period_end_time = '%s'", fiveDaysAgoRFC3339), f)

	f = WithAddress().Eq("abcd").And(WithPeriodEndTime().Lt(fiveDaysAgo)).String()
	assert.Equal(t, fmt.Sprintf("(address = 'abcd' AND period_end_time < '%s')", fiveDaysAgoRFC3339), f)

	f = WithAddress().Eq("abcd").And(WithPeriodEndTime().Lte(fiveDaysAgo)).String()
	assert.Equal(t, fmt.Sprintf("(address = 'abcd' AND period_end_time <= '%s')", fiveDaysAgoRFC3339), f)

	f = WithAddress().Eq("abcd").And(WithPeriodEndTime().Gt(fiveDaysAgo)).String()
	assert.Equal(t, fmt.Sprintf("(address = 'abcd' AND period_end_time > '%s')", fiveDaysAgoRFC3339), f)

	f = WithAddress().Eq("abcd").And(WithPeriodEndTime().Gte(fiveDaysAgo)).And(WithPeriodEndTime().Lt(now)).String()
	assert.Equal(t, fmt.Sprintf("(address = 'abcd' AND period_end_time >= '%s' AND period_end_time < '%s')", fiveDaysAgoRFC3339, nowRFC3339), f)
}
