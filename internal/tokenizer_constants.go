package internal

type LogicalOperator string

const (
	LogicalAnd LogicalOperator = "AND"
	LogicalOr  LogicalOperator = "OR"
	LogicalNot LogicalOperator = "NOT"
)

var AllLogicalOperators = []LogicalOperator{
	LogicalAnd, LogicalOr, LogicalNot,
}

func IsLogicalOperator(s string) bool {
	for _, logical := range AllLogicalOperators {
		if string(logical) == s {
			return true
		}
	}
	return false
}

func LongestLogicalOperator() int {
	longest := 0
	for _, logical := range AllLogicalOperators {
		if len(logical) > longest {
			longest = len(logical)
		}
	}
	return longest
}

// Operator represents a Meilisearch filter operator
type CompareOperator string

// Meilisearch operators constants
const (
	// Comparison operators
	OpEqual            CompareOperator = "="
	OpNotEqual         CompareOperator = "!="
	OpGreaterThanEqual CompareOperator = ">="
	OpGreaterThan      CompareOperator = ">"
	OpLessThanEqual    CompareOperator = "<="
	OpLessThan         CompareOperator = "<"
)

func AllCompareOperators() []CompareOperator {
	return []CompareOperator{
		OpEqual, OpNotEqual, OpGreaterThanEqual, OpGreaterThan, OpLessThanEqual, OpLessThan,
	}
}

func IsCompareOperator(s string) bool {
	for _, op := range AllCompareOperators() {
		if string(op) == s {
			return true
		}
	}
	return false
}

func IsCompareOperatorStart(ch rune) bool {
	return ch == '=' || ch == '!' || ch == '>' || ch == '<'
}

func LongestCompareOperator() int {
	longest := 0
	for _, op := range AllCompareOperators() {
		if len(op) > longest {
			longest = len(op)
		}
	}
	return longest
}
