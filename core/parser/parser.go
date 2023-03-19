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
	p.errors = append(p.errors, err)
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
		} else {
			p.stream.Next()
		}
	}

	return ast
}

func (p *EskemaParser) VerifySyntaxErrors() bool {

	for _, err := range p.errors {
		log.Println(err)
	}

	return len(p.errors) > 0
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

			if nextToken.Type != syntax.LiteralToken {
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
		Data: schemaDefinition,
	}
}

func (p *EskemaParser) parseType() *TypeExpression {
	typeExpression := &TypeExpression{
		Generics: make([]*TypeExpression, 0),
	}

	name := p.nextTokenMustBe(syntax.LiteralToken, syntax.PrimitiveTypeToken)

	typeExpression.Id.Name = name.Value

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

			typeExpression.Generics = append(typeExpression.Generics, genericExpr)

		}

		currentToken = p.stream.PeekCurrent()
	}

	switch currentToken.Type {
	case syntax.CommaToken,
		syntax.GreaterThanToken:
		p.stream.Next()
		break
	}

	return typeExpression
}

func (p *EskemaParser) parseAnnotation() *AnnotationExpression {
	annotationExpression := &AnnotationExpression{}
	p.nextTokenMustBe(syntax.AtToken)
	name := p.nextTokenMustBe(syntax.LiteralToken)

	annotationExpression.Type = syntax.AnnotationType(name.Value)

	p.nextTokenMustBe(syntax.ParenthesisStart)

	value := p.nextTokenMustBe(syntax.LiteralToken)
	annotationExpression.Value = value.Value

	p.nextTokenMustBe(syntax.ParenthesisEnd)

	return annotationExpression
}

func (p *EskemaParser) parseField() *FieldExpression {
	fieldExpression := &FieldExpression{}

	isAnnotation := p.stream.PeekCurrent().Type == syntax.AtToken
	fieldExpression.Annotations = make([]*AnnotationExpression, 0)
	for isAnnotation {

		annotation := p.parseAnnotation()
		fieldExpression.Annotations = append(fieldExpression.Annotations, annotation)

		isAnnotation = p.stream.PeekCurrent().Type == syntax.AtToken
	}

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
	enumDefinition := &EnumDefinition{
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
		Data: enumDefinition,
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

	return p.verifyTokenMatch(token, expectedTypes...)
}

func (p *EskemaParser) currentMustBe(expectedTypes ...syntax.TokenType) *syntax.Token {
	token := p.stream.PeekCurrent()

	return p.verifyTokenMatch(token, expectedTypes...)
}

func (p *EskemaParser) verifyTokenMatch(currentToken *syntax.Token, expectedTypes ...syntax.TokenType) *syntax.Token {

	for _, expectedType := range expectedTypes {
		if currentToken.Type == expectedType {
			return currentToken
		}
	}

	p.notifyError(
		errors.New(
			fmt.Sprintf(
				ErrUnexpectedToken,
				currentToken.Metadata,
				syntax.ToTokenTypeNiceName(expectedTypes...),
				currentToken.Value,
			),
		),
	)

	return currentToken
}

func New(stream *syntax.TokenStream) *EskemaParser {
	return &EskemaParser{
		stream: stream,
		errors: make([]error, 0),
	}
}
