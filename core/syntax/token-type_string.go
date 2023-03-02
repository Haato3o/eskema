// Code generated by "stringer -type=TokenType -output=token-type_string.go"; DO NOT EDIT.

package syntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidToken-0]
	_ = x[KeywordToken-1]
	_ = x[LiteralToken-2]
	_ = x[PrimitiveTypeToken-3]
	_ = x[WhitespaceToken-4]
	_ = x[LesserThanToken-5]
	_ = x[GreaterThanToken-6]
	_ = x[CommaToken-7]
	_ = x[ColonToken-8]
	_ = x[SemiColonToken-9]
	_ = x[QuestionMarkToken-10]
	_ = x[ScopeStart-11]
	_ = x[ScopeEnd-12]
	_ = x[NewLineToken-13]
	_ = x[EndOfFileToken-14]
}

const _TokenType_name = "InvalidTokenKeywordTokenLiteralTokenPrimitiveTypeTokenWhitespaceTokenLesserThanTokenGreaterThanTokenCommaTokenColonTokenSemiColonTokenQuestionMarkTokenScopeStartScopeEndNewLineTokenEndOfFileToken"

var _TokenType_index = [...]uint8{0, 12, 24, 36, 54, 69, 84, 100, 110, 120, 134, 151, 161, 169, 181, 195}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
