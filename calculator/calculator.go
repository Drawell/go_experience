package calculator

import (
	"fmt"
	"strconv"
	"unicode"
)

type ParseError struct {
	msg string
	idx int
}

func (p ParseError) Error() string {
	return fmt.Sprintf("%v at %v", p.msg, p.idx)
}

type Expression struct {
	text string
	ptr  int
}

func Evaluate(text string) (float64, error) {
	e := Expression{text, 0}
	value, err := e.ParseExpression()
	e.SkipSpaces()
	if err != nil {
		return 0, err
	} else if e.ptr != len(e.text) {
		return 0, ParseError{"Invalid symbols", e.ptr}
	} else {
		return value, err
	}
}

/*
Expr -> Term {('+'|'-') Term}
Term -> Factor {('*'|'/') Factor}
Factor -> Number | '(' Expr ')'.
*/

func (e *Expression) ParseExpression() (float64, error) {
	_, isOpen := e.Consume([]uint8{'('})
	lValue, err := e.ParseTerm()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.Consume([]uint8{'+', '-'}); ok; operator, ok = e.Consume([]uint8{'+', '-'}) {
		rValue, err := e.ParseTerm()
		if err != nil {
			return 0, err
		}

		lValue, err = Execute(lValue, string(operator), rValue)
		if err != nil {
			return 0, err
		}
	}

	_, isClosed := e.Consume([]uint8{')'})
	if isOpen != isClosed {
		return 0, ParseError{"There is not closing bracket", e.ptr}
	}

	return lValue, nil
}

func (e *Expression) ParseTerm() (float64, error) {
	lValue, err := e.ParseFactor()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.Consume([]uint8{'*', '/'}); ok; operator, ok = e.Consume([]uint8{'*', '/'}) {
		rValue, err := e.ParseFactor()
		if err != nil {
			return 0, err
		}

		lValue, err = Execute(lValue, string(operator), rValue)
		if err != nil {
			return 0, err
		}
	}

	return lValue, nil
}

func (e *Expression) ParseFactor() (float64, error) {
	e.SkipSpaces()
	if e.NextChar() == '(' {
		return e.ParseExpression()
	} else {
		return e.ParseOperand()
	}
}

func (e *Expression) ParseOperand() (float64, error) {
	e.SkipSpaces()
	start := e.ptr
	for ; e.ptr < len(e.text) && (unicode.IsDigit(rune(e.text[e.ptr])) || e.text[e.ptr] == '.'); e.ptr++ {
	}

	if start == e.ptr {
		return 0, ParseError{"Unable to get operand", e.ptr}
	}

	operand, _ := strconv.ParseFloat(e.text[start:e.ptr], 64)
	return operand, nil
}

func Execute(lOperand float64, operator string, rOperand float64) (float64, error) {
	switch operator {
	case "+":
		return lOperand + rOperand, nil
	case "-":
		return lOperand - rOperand, nil
	case "*":
		return lOperand * rOperand, nil
	case "/":
		if rOperand == 0 {
			return 0, ParseError{"Zero division occurs", 0}
		} else {
			return lOperand / rOperand, nil
		}
	default:
		return lOperand, nil
	}
}

func (e *Expression) NextChar() uint8 {
	e.SkipSpaces()
	return e.text[e.ptr]
}

func (e *Expression) Consume(chars []uint8) (uint8, bool) {
	e.SkipSpaces()
	if e.ptr >= len(e.text) {
		return ' ', false
	}

	for _, char := range chars {
		if e.text[e.ptr] == char {
			e.ptr++
			return char, true
		}
	}
	return ' ', false
}

func (e *Expression) SkipSpaces() {
	for ; e.ptr < len(e.text) && e.text[e.ptr] == ' '; e.ptr++ {
	}
}
