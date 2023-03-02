package main

import (
	"fmt"
	"github.com/Haato3o/eskema/core/syntax"
)

func main() {

	examplePath := "./examples/example.skm"

	lexer := syntax.NewLexer(examplePath)

	for {
		token := lexer.Lex();

		if token == nil {
			continue
		}

		fmt.Printf("[%d:%d] %v : %s\n", token.Metadata.Line, token.Metadata.Column, token.Type, token.Value)

		if token.Type == syntax.EndOfFileToken {
			break
		}
	}

}
