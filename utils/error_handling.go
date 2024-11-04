package utils

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

func (e *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("Syntax error at line %d:%d - %s\n", line, column, msg)
}
