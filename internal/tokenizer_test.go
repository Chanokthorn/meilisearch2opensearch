package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeSimpleEquality(t *testing.T) {
	input := `subcategory.slug = "seo"`
	expected := []Token{
		{TokenType: TokenTypeField, Value: "subcategory.slug"},
		{TokenType: TokenTypeOperator, Value: "="},
		{TokenType: TokenTypeValue, Value: "seo"},
	}

	tokens, err := Tokenize(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, tokens)
}

func TestTokenizeWithLogicalOperator(t *testing.T) {
	input := `subcategory.slug = "seo" AND rating >= 3`
	expected := []Token{
		{TokenType: TokenTypeField, Value: "subcategory.slug"},
		{TokenType: TokenTypeOperator, Value: "="},
		{TokenType: TokenTypeValue, Value: "seo"},
		{TokenType: TokenTypeLogicalOperator, Value: "AND"}, // if logicals are treated as fields
		{TokenType: TokenTypeField, Value: "rating"},
		{TokenType: TokenTypeOperator, Value: ">="},
		{TokenType: TokenTypeValue, Value: "3"},
	}

	tokens, err := Tokenize(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, tokens)
}

func TestTokenizeParenthesisAndMultipleExpressions(t *testing.T) {
	input := `(subcategory.slug = "seo") AND (base_price >= 100)`
	expected := []Token{
		{TokenType: TokenTypeParanthesisOpen, Value: "("},
		{TokenType: TokenTypeField, Value: "subcategory.slug"},
		{TokenType: TokenTypeOperator, Value: "="},
		{TokenType: TokenTypeValue, Value: "seo"},
		{TokenType: TokenTypeParanthesisClose, Value: ")"},
		{TokenType: TokenTypeLogicalOperator, Value: "AND"}, // same caveat as above
		{TokenType: TokenTypeParanthesisOpen, Value: "("},
		{TokenType: TokenTypeField, Value: "base_price"},
		{TokenType: TokenTypeOperator, Value: ">="},
		{TokenType: TokenTypeValue, Value: "100"},
		{TokenType: TokenTypeParanthesisClose, Value: ")"},
	}

	tokens, err := Tokenize(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, tokens)
}