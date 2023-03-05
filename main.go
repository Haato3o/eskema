package main

import (
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/core/syntax"
	"github.com/Haato3o/eskema/core/visualization"
	"github.com/Haato3o/eskema/emitter/languages"
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

	emitter := languages.NewCSharpEmitter()

	ktCode := emitter.Emit(ast)

	println(ktCode)
}
