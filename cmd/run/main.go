package main

import (
	"strings"
)

// Operator represents a Meilisearch filter operator
type Operator string

// Meilisearch operators constants
const (
	// Comparison operators
	OpEqual            Operator = "="
	OpNotEqual         Operator = "!="
	OpGreaterThanEqual Operator = ">="
	OpGreaterThan      Operator = ">"
	OpLessThanEqual    Operator = "<="
	OpLessThan         Operator = "<"

	// Set operators
	OpIn    Operator = "IN"
	OpNotIn Operator = "NOT IN"

	// Range operator
	OpTo Operator = "TO"

	// Existence operators
	OpExists    Operator = "EXISTS"
	OpNotExists Operator = "NOT EXISTS"

	// Null operators
	OpIsNull    Operator = "IS NULL"
	OpIsNotNull Operator = "IS NOT NULL"

	// Empty operators
	OpIsEmpty    Operator = "IS EMPTY"
	OpIsNotEmpty Operator = "IS NOT EMPTY"

	// String operators
	OpContains      Operator = "CONTAINS"
	OpNotContains   Operator = "NOT CONTAINS"
	OpStartsWith    Operator = "STARTS WITH"
	OpNotStartsWith Operator = "NOT STARTS WITH"

	// Geo operators
	OpGeoRadius      Operator = "_geoRadius"
	OpGeoBoundingBox Operator = "_geoBoundingBox"
)

// AllOperators contains all Meilisearch operators
var AllOperators = []Operator{
	OpEqual, OpNotEqual, OpGreaterThanEqual, OpGreaterThan, OpLessThanEqual, OpLessThan,
	OpIn, OpNotIn, OpTo, OpExists, OpNotExists, OpIsNull, OpIsNotNull,
	OpIsEmpty, OpIsNotEmpty, OpContains, OpNotContains, OpStartsWith, OpNotStartsWith,
	OpGeoRadius, OpGeoBoundingBox,
}

// IsOperator checks if a string is a valid Meilisearch operator
func IsOperator(s string) bool {
	for _, op := range AllOperators {
		if strings.ToUpper(string(op)) == s {
			return true
		}
	}
	return false
}

// GetOperatorType categorizes the operator type
func GetOperatorType(op Operator) string {
	switch op {
	case OpEqual, OpNotEqual, OpGreaterThanEqual, OpGreaterThan, OpLessThanEqual, OpLessThan:
		return "comparison"
	case OpIn, OpNotIn:
		return "set"
	case OpTo:
		return "range"
	case OpExists, OpNotExists:
		return "existence"
	case OpIsNull, OpIsNotNull:
		return "null"
	case OpIsEmpty, OpIsNotEmpty:
		return "empty"
	case OpContains, OpNotContains, OpStartsWith, OpNotStartsWith:
		return "string"
	case OpGeoRadius, OpGeoBoundingBox:
		return "geo"
	default:
		return "unknown"
	}
}

type Logical string

// Logical constants
const (
	LogicalAnd Logical = "AND"
	LogicalOr  Logical = "OR"
	LogicalNot Logical = "NOT"
)

var AllLogical = []Logical{
	LogicalAnd, LogicalOr, LogicalNot,
}

func IsLogical(s string) bool {
	for _, logical := range AllLogical {
		if string(logical) == s {
			return true
		}
	}
	return false
}

type TokenType int

// TokenType constants
const (
	TokenTypeOperator TokenType = iota
	TokenTypeValue
	TokenTypeField
	TokenTypeParnOpen
	TokenTypeParnClose
)

type Token struct {
	TokenType string
	Value     string
}

type TokenizerMode int 
const (
	ModeDefault TokenizerMode = iota
	ModeString
)

/*
for each character:
inital mode to ModeDefault

if mode default:
	if it's spcae
		skip
	if it's double quote
		switch mode to ModeString
	if it's parenthesis open
		append Token{TokenType: TokenTypeParnOpen, Value: "("} to tokens
		continue
	if it's parenthesis close
		append Token{TokenType: TokenTypeParnClose, Value: ")"} to tokens
		continue
	if it's a string character (alphanumeric or underscore)
		append to currentToken
		if currentToken is not empty -> can be operator or logical
			get TokenType based on currentToken
			append Token{TokenType: TokenTypeField or TokenTypeValue, Value: currentToken} to tokens
			currentToken = ""

			if is operator // case field = base_price
				switch mode to ModeString
			
		continue

	if mode is ModeString:
		if it's escape character
			append Token{TokenType: TokenTypeValue, Value: currentToken} to tokens
			switch mode to ModeStringEscape	
			continue
		if it's double quote
			switch mode to ModeDefault
			if currentToken is not empty
				get TokenType based on currentToken
				append Token{TokenType: TokenTypeField or TokenTypeValue, Value: currentToken} to tokens
				currentToken = ""
			continue
		if EOF
			error if currentToken is not empty

	if mode is ModeStringEscape:
		if it's double quote
			append to currentToken
			switch mode to ModeString
			continue
		if it's any other character
			append `\<char>` to currentToken
			switch mode to ModeString
			continue


*/
func Tokenize(input string) []Token {
	var tokens []Token
	// var currentToken string
	// var mode 




	return tokens
}

func main() {
}
