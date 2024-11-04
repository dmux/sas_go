package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"log"
	"os"
	"sas_go/interpreter"
	"sas_go/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to a .sas file as an argument.")
	}

	sasFilePath := os.Args[1]

	content, err := ioutil.ReadFile(sasFilePath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", sasFilePath, err)
	}

	input := antlr.NewInputStream(string(content))

	lexer := parser.NewSASLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSASParser(tokens)

	listener := interpreter.NewSASListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	fmt.Println("Interpretation completed.")
}
