package main

import (
	"bytes"
	"encoding/json"
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/core/syntax"
	"os"
)

func main() {

	examplePath := "./examples/example.skm"

	lexer := syntax.NewLexer(examplePath)
	tokens := lexer.Lex()

	eskemaParser := parser.New(tokens)

	ast := eskemaParser.Parse()

	eskemaParser.VerifySyntaxErrors()

	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(ast)
	os.WriteFile("output.json", buffer.Bytes(), 0644)
}
