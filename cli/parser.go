package cli

import "flag"

func ParseArguments() *EskemaArguments {
	fileName := flag.String("filename", "", "Path to the eskema file")
	language := flag.String("language", "", "Language to parse the schema to")
	shouldPrintAst := flag.Bool("ast", false, "Whether the generated AST should be displayed or not")
	shouldPrintSupportedLanguages := flag.Bool("langs", false, "Use this command to display the supported languages for Eskema")

	flag.Parse()

	return &EskemaArguments{
		FileName:                      *fileName,
		Language:                      *language,
		ShouldPrintAST:                *shouldPrintAst,
		ShouldPrintSupportedLanguages: *shouldPrintSupportedLanguages,
	}
}
