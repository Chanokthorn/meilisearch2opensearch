package internal

type NodeType int

const (
	NodeTypeField NodeType = iota
	NodeTypeValue
	NodeTypeCompareOperator
	NodeTypeLogicalOperator
)

type Node struct {
	Type NodeType
	Value string
	Children []*Node
}

/*
BuildAST constructs an abstract syntax tree (AST) from a given input string.
f([]token) -> Node

common cases:
if token is open parenthesis:
if token is logical operator
  
if first token is a field, then next should be an operator, then a value
if next token is a logical operator, then it should be followed by another field or parenthesis
if next token is a close parenthesis, then it should close the last open parenthesis

base cases:


*/
func BuildAST(tokens []Token) (*Node, error) {
	
}
