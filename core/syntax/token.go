package syntax

import "fmt"

//go:generate stringer -type=TokenType -output=token-type_string.go

type TokenType int

const (
	InvalidToken TokenType = iota

	KeywordToken
	LiteralToken
	PrimitiveTypeToken

	WhitespaceToken
	LesserThanToken
	GreaterThanToken
	CommaToken
	ColonToken
	SemiColonToken
	QuestionMarkToken

	ScopeStartToken
	ScopeEndToken

	NewLineToken
	EndOfFileToken
)

type Metadata struct {
	Filename string
	Offset   int64
	Line     int64
	Column   int64
}

func (i Metadata) String() string {
	return fmt.Sprintf("%s [%d:%d]", i.Filename, i.Line, i.Column)
}

type Token struct {
	Metadata *Metadata
	Value    string
	Type     TokenType
}
