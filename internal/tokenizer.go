package internal

import "fmt"

type TokenType int

const (
	TokenTypeParanthesisOpen TokenType = iota
	TokenTypeParanthesisClose
	TokenTypeOperator
	TokenTypeLogicalOperator
	TokenTypeValue
	TokenTypeField
)

type Token struct {
	TokenType TokenType
	Value     string
}

type TokenizerMode int

const (
	ModeDefault TokenizerMode = iota
	ModeField
	ModeOperator
	ModeValue
	ModeSeeking
)

func Tokenize(input string) ([]Token, error) {
	// set up initial states
	var tokens []Token
	var mode TokenizerMode = ModeDefault
	var currentBuffer string
	var seekingFor TokenizerMode

	// flag used in ModeField and ModeValue
	var isEscapeNext bool
	var isWithinText bool
	var isDoubleQuoteOpen bool

	// iterate through each character
	for i := 0; i < len(input); {
		char := rune(input[i])
		switch mode {
		case ModeSeeking:
			if char == ' ' {
				i++
				continue // skip spaces
			}
			mode = seekingFor
			continue
		case ModeDefault:
			if char == ' ' {
				i++
				continue // skip spaces
			}
			if char == '(' {
				tokens = append(tokens, Token{TokenType: TokenTypeParanthesisOpen, Value: "("})
				i++
				continue
			}
			if char == ')' {
				tokens = append(tokens, Token{TokenType: TokenTypeParanthesisClose, Value: ")"})
				i++
				continue
			}
			if char == '"' {
				isDoubleQuoteOpen = true
				mode = ModeField // switch to field mode
				i++
				continue
			}
			// greedily check for logical operators
			logicalOperatorFound := false
			for j := i; j < min(i+LongestLogicalOperator(), len(input)); j++ {
				currentLogicalCheck := input[i : j+1]
				if IsLogicalOperator(currentLogicalCheck) {
					logicalOperatorFound = true
					tokens = append(tokens, Token{TokenType: TokenTypeLogicalOperator, Value: currentLogicalCheck})
					i += len(currentLogicalCheck)
					currentBuffer = ""
					mode = ModeValue // switch to value mode
					continue
				}
			}
			if logicalOperatorFound {
				continue
			}
			// greedily check for compare operators
			compareOperatorFound := false
			for j := i; j < min(i+LongestCompareOperator(), len(input)); j++ {
				currentCompareCheck := input[i : j+1]
				if IsCompareOperator(currentCompareCheck) {
					tokens = append(tokens, Token{TokenType: TokenTypeOperator, Value: currentCompareCheck})
					i += len(currentCompareCheck)
					currentBuffer = ""
					mode = ModeValue // switch to value mode
					compareOperatorFound = true
					break
				}
			}
			if compareOperatorFound {
				mode = ModeSeeking
				seekingFor = ModeValue // switch to seeking for value mode
				continue
			}
			// if no logical or compare operator found, treat as field
			mode = ModeField
			continue
		case ModeField:
			if char == '\\' && isDoubleQuoteOpen {
				isEscapeNext = true
				i++
				continue
			}
			// handle immediate operators
			if IsCompareOperatorStart(char) && !isDoubleQuoteOpen {
				if currentBuffer != "" {
					tokens = append(tokens, Token{TokenType: TokenTypeField, Value: currentBuffer})
					currentBuffer = ""
				}
				mode = ModeOperator // switch to operator mode without adding to buffer
				continue
			}
			if char == '"' {
				if isEscapeNext {
					currentBuffer += string(char)
					isEscapeNext = false
					continue
				}
				// end of field value, append token and reset state
				tokens = append(tokens, Token{
					TokenType: TokenTypeField,
					Value:     currentBuffer,
				})
				currentBuffer = ""
				isDoubleQuoteOpen = false
				mode = ModeSeeking
				seekingFor = ModeValue // switch to seeking for value mode
				i++
				continue
			}
			// is within a field value
			if isDoubleQuoteOpen || char != ' ' {
				currentBuffer += string(char)
				i++
				continue
			}
			// end of field, append token and reset state
			if currentBuffer != "" {
				tokens = append(tokens, Token{TokenType: TokenTypeField, Value: currentBuffer})
				currentBuffer = ""
			}
			mode = ModeSeeking
			seekingFor = ModeOperator
			i++
			continue

		case ModeValue:
			if char == '\\' && isDoubleQuoteOpen {
				isEscapeNext = true
				i++
				continue
			}
			if char == '"' {
				if !isWithinText { // start of a value
					isWithinText = true
					isDoubleQuoteOpen = true
					currentBuffer = "" // reset buffer for new value
					i++
					continue
				}
				if isEscapeNext {
					currentBuffer += string(char)
					isEscapeNext = false
					continue
				}
				// end of value, append token and reset state
				tokens = append(tokens, Token{
					TokenType: TokenTypeValue,
					Value:     currentBuffer,
				})
				currentBuffer = ""
				isDoubleQuoteOpen = false
				mode = ModeDefault // switch back to default mode
				i++
				continue
			}
			// if is parenthesis, switch to default mode
			if char == '(' || char == ')' {
				if currentBuffer != "" {
					tokens = append(tokens, Token{TokenType: TokenTypeValue, Value: currentBuffer})
					currentBuffer = ""
					isWithinText = false
				}
				mode = ModeDefault // switch back to default mode
				continue
			}
			// is within a text
			if isDoubleQuoteOpen || char != ' ' {
				currentBuffer += string(char)
				i++
				continue
			}
			// end of value, append token and reset state
			if currentBuffer != "" {
				tokens = append(tokens, Token{TokenType: TokenTypeValue, Value: currentBuffer})
				currentBuffer = ""
				isWithinText = false
			}
			i++
			mode = ModeDefault // switch back to default mode
			continue
		// ModeOperator greedily checks for operators
		case ModeOperator:
			// check lookahead for multi-character operators; currently max is 2 characters
			if i+1 < len(input) {
				currentOperatorString := input[i : i+2]
				if IsCompareOperator(currentOperatorString) {
					tokens = append(tokens, Token{TokenType: TokenTypeOperator, Value: currentOperatorString})
					i += 2
					currentBuffer = ""
					mode = ModeSeeking
					seekingFor = ModeValue // switch to seeking for value mode
					continue
				}
			}
			// if single character operator
			if IsCompareOperator(string(char)) {
				tokens = append(tokens, Token{TokenType: TokenTypeOperator, Value: string(char)})
				i++
				currentBuffer = ""
				mode = ModeSeeking
				seekingFor = ModeValue // switch to seeking for value mode
				continue
			}
			// if not a valid operator, throw error
			return nil, fmt.Errorf("unexpected token: %s for operator", string(char))
		}
	}

	// handle remaining token if exists
	if currentBuffer != "" {
		if mode != ModeValue {
			return nil, fmt.Errorf("unexpected end of input, expected value for token: %s", currentBuffer)
		}
		if isDoubleQuoteOpen {
			return nil, fmt.Errorf("unclosed double quote in input: %s", currentBuffer)
		}
		tokens = append(tokens, Token{TokenType: TokenTypeValue, Value: currentBuffer})
	}

	return tokens, nil
}
