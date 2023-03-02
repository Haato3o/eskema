package parser

import (
	"errors"
	"fmt"
	"github.com/Haato3o/eskema/core/syntax"
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
		token := p.stream.Current()

		if token.Type == syntax.EndOfFileToken {
			break
		}

		if token.Type == syntax.KeywordToken {
			expr := p.parseKeyword()
			ast.Expr = append(ast.Expr, expr)
		} else {
			p.stream.Next()
		}
	}

	return ast
}

func (p *EskemaParser) parseKeyword() *EskemaExpression {
	token := p.stream.Next()

	_, keywordType := syntax.IsKeyword(token.Value)

	if keywordType == syntax.EnumKeyword {
		enumDefinition := EnumDefinition{
			Values: make([]string, 0),
		}
		name := p.nextTokenMustBe(syntax.LiteralToken)

		enumDefinition.Id.Name = name.Value

		p.nextTokenMustBe(syntax.ScopeStart)

		for {
			enumValue := p.parseEnumValue()

			if enumValue == nil {
				break
			}

			enumDefinition.Values = append(enumDefinition.Values, enumValue.Value)
		}

		p.nextTokenMustBe(syntax.ScopeEnd)
		p.nextTokenMustBe(syntax.SemiColonToken)

		return &EskemaExpression{
			Type: EnumExpr,
			Data: &enumDefinition,
		}
	}

	return nil
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

func (p *EskemaParser) nextTokenMustBe(expectedType syntax.TokenType) *syntax.Token {
	token := p.stream.Next()

	if token.Type != expectedType {

		p.notifyError(errors.New(fmt.Sprintf(ErrUnexpectedToken, expectedType, token.Type)))

		return nil
	}

	return token
}

func New(stream *syntax.TokenStream) *EskemaParser {
	return &EskemaParser{
		stream: stream,
		errors: make([]error, 0),
	}
}
