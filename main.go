package main

import (
	"bytes"
	"encoding/json"
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/core/syntax"
	"github.com/Haato3o/eskema/core/visualization"
	"os"
)

func main() {

	examplePath := "./examples/example.skm"

	lexer := syntax.NewLexer(examplePath)
	tokens := lexer.Lex()

	eskemaParser := parser.New(tokens)

	ast := eskemaParser.Parse()

	if hasErrors := eskemaParser.VerifySyntaxErrors(); hasErrors {
		return
	}

	visualization.VisualizeTree(ast)

	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(ast)
	os.WriteFile("output.json", buffer.Bytes(), 0644)
}
