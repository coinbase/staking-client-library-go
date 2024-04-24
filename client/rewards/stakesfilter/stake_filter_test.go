package stakesfilter_test

import (
	"testing"
	"time"

	filter "github.com/coinbase/staking-client-library-go/client/rewards/stakesfilter"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/suite"
)

func TestListStakesFilterSuite(t *testing.T) {
	suite.Run(t, new(ListStakesFilterSuite))
}

type ListStakesFilterSuite struct {
	suite.Suite
}

func (s *ListStakesFilterSuite) TestListStakesFilter_WithAddress() {
	f := filter.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListStakesFilterSuite) TestListStakesFilter_EvaluationTime() {
	var f string

	specificTime := time.Date(2024, time.February, 1, 0, 0, 1, 0, time.UTC)

	f = filter.WithEvaluationTime().Eq(specificTime).String()
	assert.Equal(s.T(), "evaluation_time = '2024-02-01T00:00:01Z'", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Lt(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time < '2024-02-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Lte(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time <= '2024-02-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Gt(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time > '2024-02-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Gte(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time >= '2024-02-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Neq(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time != '2024-02-01T00:00:01Z')", f)
}

func (s *ListStakesFilterSuite) TestListStakesFilter_EvaluationTime_Double() {
	var f string

	specificTime1 := time.Date(2024, time.February, 1, 0, 0, 1, 0, time.UTC)
	specificTime2 := time.Date(2024, time.March, 1, 0, 0, 1, 0, time.UTC)

	f = filter.WithEvaluationTime().Eq(specificTime1).And(filter.WithEvaluationTime().Neq(specificTime2)).String()
	assert.Equal(s.T(), "(evaluation_time = '2024-02-01T00:00:01Z' AND evaluation_time != '2024-03-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Gt(specificTime1)).And(filter.WithEvaluationTime().Lt(specificTime2)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time > '2024-02-01T00:00:01Z' AND evaluation_time < '2024-03-01T00:00:01Z')", f)

	f = filter.WithAddress().Eq("abcd").And(filter.WithEvaluationTime().Gte(specificTime1)).And(filter.WithEvaluationTime().Lte(specificTime2)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time >= '2024-02-01T00:00:01Z' AND evaluation_time <= '2024-03-01T00:00:01Z')", f)
}
