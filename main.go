package main

import (
	"github.com/Haato3o/eskema/cli"
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/core/syntax"
	"github.com/Haato3o/eskema/core/visualization"
	"log"
	"os"
)

func main() {

	args := cli.ParseArguments()

	if args.ShouldPrintSupportedLanguages {
		cli.PrintSupportedLanguages()
		return
	}

	if err := args.VerifyRequired(); err != nil {
		log.Fatalln(err)
	}

	lexer := syntax.NewLexerFromFile(args.FileName)
	tokens := lexer.Lex()
	eskemaParser := parser.New(tokens)
	ast := eskemaParser.Parse()

	if hasErrors := eskemaParser.VerifySyntaxErrors(); hasErrors {
		return
	}

	if args.ShouldPrintAST {
		visualization.VisualizeTree(ast)
	}

	emitter, err := cli.GetLanguageEmitter(args.Language)

	if err != nil {
		log.Fatalln(err)
	}

	code := emitter.Emit(ast)

	if args.Output != "" {
		file, _ := os.OpenFile(args.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

		_, _ = file.WriteString(code)

		_ = file.Close()
	} else {
		println(code)
	}

}
