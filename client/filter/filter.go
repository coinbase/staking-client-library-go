package filter

import (
	"fmt"
	"strings"
	"time"
)

// Supported operations taken from https://google.aip.dev/160

type ComparisonOperation string

const (
	Equals        ComparisonOperation = "="
	NotEquals     ComparisonOperation = "!="
	LessThan      ComparisonOperation = "<"
	LessEquals    ComparisonOperation = "<="
	GreaterEquals ComparisonOperation = ">="
	GreaterThan   ComparisonOperation = ">"
)

type LogicalOperation string

const (
	And LogicalOperation = "AND"
)

// Filterer interface for both Term and LogicalExpressions.
type Filterer interface {
	And(other Filterer) *LogicalExpressions
	String() string
}

// LogicalExpression represents a pair of filter terms joined with a logical operator.
type LogicalExpression struct {
	Op     LogicalOperation
	Filter Filterer
}

// LogicalExpressions represents a list of logical expressions.
type LogicalExpressions struct {
	expressions []LogicalExpression
}

// Term helps represent a single filter condition, e.g. "field = value".
type Term struct {
	field string
	op    ComparisonOperation
	value interface{}
}

func (t *Term) And(other Filterer) *LogicalExpressions {
	return &LogicalExpressions{
		expressions: []LogicalExpression{
			{Op: And, Filter: t},
			{Op: And, Filter: other},
		},
	}
}

func (t *Term) String() string {
	parsedValue := t.value

	if timeValue, ok := t.value.(time.Time); ok {
		// If the type assertion is successful, the value is a time.Time,
		// so we need to parse it to a string in RFC3339 format.
		parsedValue = fmt.Sprintf("'%s'", timeValue.Format(time.RFC3339))
	} else if intValue, ok := t.value.(int); ok {
		// If the type assertion is successful, the value is an int
		parsedValue = fmt.Sprintf("%d", intValue)
	} else if stringValue, ok := t.value.(string); ok {
		// If the type assertion is successful, the value is a string
		parsedValue = fmt.Sprintf("'%s'", stringValue)
	}

	return fmt.Sprintf("%s %s %s", t.field, t.op, parsedValue)
}

func NewTerm(field string, op ComparisonOperation, value interface{}) *Term {
	return &Term{field: field, op: op, value: value}
}

func (le *LogicalExpressions) And(other Filterer) *LogicalExpressions {
	le.expressions = append(le.expressions, LogicalExpression{Op: And, Filter: other})

	return le
}

func (le *LogicalExpressions) String() string {
	if len(le.expressions) == 0 {
		return ""
	}

	parts := make([]string, len(le.expressions)*2-1)

	index := 0

	for i, expr := range le.expressions {
		if i > 0 {
			parts[index] = fmt.Sprintf(" %s ", expr.Op)
			index++
		}

		parts[index] = expr.Filter.String()
		index++
	}

	return "(" + strings.Join(parts, "") + ")"
}
