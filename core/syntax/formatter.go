package syntax

import "fmt"

var equivalents = map[TokenType]string{
	WhitespaceToken:   " ",
	LesserThanToken:   "<",
	GreaterThanToken:  ">",
	CommaToken:        ",",
	ColonToken:        ":",
	SemiColonToken:    ";",
	QuestionMarkToken: "?",
	ScopeStartToken:   "{",
	ScopeEndToken:     "}",

	LiteralToken: "Literal",

	EndOfFileToken: "EOF",
}

func ToTokenTypeNiceName(tokens ...TokenType) string {
	formatted := ""

	for _, token := range tokens {
		equivalent, exists := equivalents[token]

		if !exists {
			equivalent = token.String()
		}

		if formatted == "" {
			formatted += fmt.Sprintf("'%s'", equivalent)
		} else {
			formatted += fmt.Sprintf(" or '%s'", equivalent)
		}
	}

	return formatted
}
