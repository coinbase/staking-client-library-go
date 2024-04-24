package stakes_test

import (
	"testing"
	"time"

	"github.com/coinbase/staking-client-library-go/client/rewards/stakes"
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
	f := stakes.WithAddress().Eq("abcd").String()
	assert.Equal(s.T(), "address = 'abcd'", f)
}

func (s *ListStakesFilterSuite) TestListStakesFilter_EvaluationTime() {
	var f string

	specificTime := time.Date(2024, time.February, 1, 0, 0, 1, 0, time.UTC)

	f = stakes.WithEvaluationTime().Eq(specificTime).String()
	assert.Equal(s.T(), "evaluation_time = '2024-02-01T00:00:01Z'", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Lt(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time < '2024-02-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Lte(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time <= '2024-02-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Gt(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time > '2024-02-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Gte(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time >= '2024-02-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Neq(specificTime)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time != '2024-02-01T00:00:01Z')", f)
}

func (s *ListStakesFilterSuite) TestListStakesFilter_EvaluationTime_Double() {
	var f string

	specificTime1 := time.Date(2024, time.February, 1, 0, 0, 1, 0, time.UTC)
	specificTime2 := time.Date(2024, time.March, 1, 0, 0, 1, 0, time.UTC)

	f = stakes.WithEvaluationTime().Eq(specificTime1).And(stakes.WithEvaluationTime().Neq(specificTime2)).String()
	assert.Equal(s.T(), "(evaluation_time = '2024-02-01T00:00:01Z' AND evaluation_time != '2024-03-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Gt(specificTime1)).And(stakes.WithEvaluationTime().Lt(specificTime2)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time > '2024-02-01T00:00:01Z' AND evaluation_time < '2024-03-01T00:00:01Z')", f)

	f = stakes.WithAddress().Eq("abcd").And(stakes.WithEvaluationTime().Gte(specificTime1)).And(stakes.WithEvaluationTime().Lte(specificTime2)).String()
	assert.Equal(s.T(), "(address = 'abcd' AND evaluation_time >= '2024-02-01T00:00:01Z' AND evaluation_time <= '2024-03-01T00:00:01Z')", f)
}
