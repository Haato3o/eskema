package syntax

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

	ScopeStart
	ScopeEnd

	NewLineToken
	EndOfFileToken
)

type Metadata struct {
	Filename string
	Offset   int64
	Line     int64
	Column   int64
}

type Token struct {
	Metadata *Metadata
	Value    string
	Type     TokenType
}
