package tests

import (
	"github.com/antlr4-go/antlr/v4"
	"sas_go/interpreter"
	"sas_go/parser"
	"testing"
)

func TestParser(t *testing.T) {
	input := antlr.NewInputStream(`
        DATA _NULL_;
		PUT 'Hello, World!';
		RUN;
    `)

	lexer := parser.NewSASLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSASParser(tokens)

	listener := interpreter.NewSASListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
}
