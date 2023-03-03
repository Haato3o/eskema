package parser

import (
	"errors"
	"fmt"
	"github.com/Haato3o/eskema/core/syntax"
	"log"
)

type EskemaParser struct {
	stream *syntax.TokenStream
	errors []error
}

func (p *EskemaParser) notifyError(err error) {
	log.Fatal(err)

	//p.errors = append(p.errors, err)
}

func (p *EskemaParser) Parse() *EskemaTree {
	ast := &EskemaTree{
		Expr: make([]*EskemaExpression, 0),
	}

	for {
		token := p.currentMustBe(syntax.KeywordToken, syntax.EndOfFileToken)

		if token.Type == syntax.KeywordToken {
			expr := p.parseKeyword()
			ast.Expr = append(ast.Expr, expr)
		} else if token.Type == syntax.EndOfFileToken {
			break
		}
	}

	return ast
}

func (p *EskemaParser) parseKeyword() *EskemaExpression {
	token := p.stream.Next()

	_, keywordType := syntax.IsKeyword(token.Value)

	switch keywordType {
	case syntax.EnumKeyword:
		return p.parseEnum()
	case syntax.SchemaKeyword:
		return p.parseSchema()
	default:
		return nil
	}
}

func (p *EskemaParser) parseSchema() *EskemaExpression {
	schemaDefinition := &SchemaDefinition{
		Fields: make([]*FieldExpression, 0),
	}
	name := p.nextTokenMustBe(syntax.LiteralToken)

	schemaDefinition.Id.Name = name.Value

	currentToken := p.stream.PeekCurrent()

	if currentToken.Type == syntax.LesserThanToken {
		p.stream.Next()
		for {
			nextToken := p.stream.PeekCurrent()

			if nextToken.Type != syntax.LiteralToken &&
				nextToken.Type != syntax.PrimitiveTypeToken {
				break
			}

			genericExpr := p.parseType()

			if genericExpr == nil {
				break
			}

			schemaDefinition.Generics = append(schemaDefinition.Generics, genericExpr)
		}
	}

	currentToken = p.nextTokenMustBe(syntax.ScopeStartToken)

	for {
		currentToken = p.stream.PeekCurrent()

		if currentToken.Type == syntax.ScopeEndToken {
			break
		}

		fieldExpr := p.parseField()

		if fieldExpr == nil {
			break
		}

		schemaDefinition.Fields = append(schemaDefinition.Fields, fieldExpr)
	}

	p.nextTokenMustBe(syntax.ScopeEndToken)
	p.nextTokenMustBe(syntax.SemiColonToken)

	return &EskemaExpression{
		Type: SchemaExpr,
		Data: &schemaDefinition,
	}
}

func (p *EskemaParser) parseType() *TypeExpression {
	typeExpression := &TypeExpression{
		Generics: make([]*TypeExpression, 0),
	}

	name := p.nextTokenMustBe(syntax.LiteralToken, syntax.PrimitiveTypeToken)

	typeExpression.Id.Name = name.Value

	currentToken := p.stream.PeekCurrent()

	switch currentToken.Type {
	case syntax.CommaToken,
		syntax.GreaterThanToken:
		p.stream.Next()
		break
	case syntax.LesserThanToken:
		p.stream.Next()

		for {
			nextToken := p.stream.PeekCurrent()

			if nextToken.Type != syntax.LiteralToken &&
				nextToken.Type != syntax.PrimitiveTypeToken {
				break
			}

			genericExpr := p.parseType()

			if genericExpr == nil {
				break
			}

			typeExpression.Generics = append(typeExpression.Generics, genericExpr)

		}
		break
	}

	return typeExpression
}

func (p *EskemaParser) parseField() *FieldExpression {
	fieldExpression := &FieldExpression{}

	name := p.nextTokenMustBe(syntax.LiteralToken)

	fieldExpression.Id.Name = name.Value

	p.nextTokenMustBe(syntax.ColonToken)

	fieldExpression.Type = p.parseType()

	endToken := p.stream.PeekCurrent()

	if endToken.Type == syntax.QuestionMarkToken {
		p.stream.Next()

		fieldExpression.IsOptional = true

		endToken = p.stream.PeekCurrent()
	}

	if endToken.Type == syntax.CommaToken {
		p.stream.Next()
	}

	return fieldExpression
}

func (p *EskemaParser) parseEnum() *EskemaExpression {
	enumDefinition := EnumDefinition{
		Values: make([]string, 0),
	}
	name := p.nextTokenMustBe(syntax.LiteralToken)

	enumDefinition.Id.Name = name.Value

	p.nextTokenMustBe(syntax.ScopeStartToken)

	for {
		enumValue := p.parseEnumValue()

		if enumValue == nil {
			break
		}

		enumDefinition.Values = append(enumDefinition.Values, enumValue.Value)
	}

	p.nextTokenMustBe(syntax.ScopeEndToken)
	p.nextTokenMustBe(syntax.SemiColonToken)

	return &EskemaExpression{
		Type: EnumExpr,
		Data: &enumDefinition,
	}
}

func (p *EskemaParser) parseEnumValue() *syntax.Token {
	next := p.stream.Peek()
	current := p.stream.Next()

	if current.Type != syntax.LiteralToken {
		p.stream.Prev()
		return nil
	}

	if next.Type == syntax.CommaToken {
		p.stream.Next()
	}

	return current
}

func (p *EskemaParser) nextTokenMustBe(expectedTypes ...syntax.TokenType) *syntax.Token {
	token := p.stream.Next()

	for _, expectedType := range expectedTypes {
		if token.Type == expectedType {
			return token
		}
	}

	p.notifyError(errors.New(fmt.Sprintf(ErrUnexpectedToken, token.Metadata, expectedTypes, token.Type)))

	return nil
}

func (p *EskemaParser) currentMustBe(expectedTypes ...syntax.TokenType) *syntax.Token {
	token := p.stream.PeekCurrent()

	for _, expectedType := range expectedTypes {
		if token.Type == expectedType {
			return token
		}
	}

	p.notifyError(errors.New(fmt.Sprintf(ErrUnexpectedToken, token.Metadata, expectedTypes, token.Type)))

	return nil
}

func New(stream *syntax.TokenStream) *EskemaParser {
	return &EskemaParser{
		stream: stream,
		errors: make([]error, 0),
	}
}
