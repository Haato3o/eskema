package syntax

var keywords = map[string]Keyword{
	"schema": SchemaKeyword,
	"enum":   EnumKeyword,
}

var primitives = map[string]Primitive{
	"String":    String,
	"Char":      Char,
	"UInt8":     UInt8,
	"UInt16":    UInt16,
	"UInt32":    UInt32,
	"UInt64":    UInt64,
	"Int8":      Int8,
	"Int16":     Int16,
	"Int32":     Int32,
	"Int64":     Int64,
	"Float":     Float,
	"Double":    Double,
	"TimeStamp": TimeStamp,
	"Date":      Date,
	"DateTime":  DateTime,
	"Array":     Array,
	"Map":       Map,
}

var tokens = map[byte]TokenType{
	'\x00': InvalidToken,
	'\t':   InvalidToken,
	'\r':   InvalidToken,

	' ': WhitespaceToken,
	'<': LesserThanToken,
	'>': GreaterThanToken,
	',': CommaToken,
	':': ColonToken,
	';': SemiColonToken,
	'?': QuestionMarkToken,

	'{': ScopeStartToken,
	'}': ScopeEndToken,

	'\n': NewLineToken,
}

func IsSpecialToken(value byte) (bool, TokenType) {
	token, exists := tokens[value]

	return exists, token
}

func IsKeyword(value string) (bool, Keyword) {
	keyword, exists := keywords[value]

	return exists, keyword
}

func IsPrimitiveType(value string) (bool, Primitive) {
	primitive, exists := primitives[value]

	return exists, primitive
}
