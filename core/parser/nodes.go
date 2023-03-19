package parser

import "github.com/Haato3o/eskema/core/syntax"

type EskemaExprType int

const (
	SchemaExpr EskemaExprType = iota
	EnumExpr
)

type EskemaExpression struct {
	Type EskemaExprType
	Data interface{}
}

type IdentifierExpression struct {
	Name string
}

type SchemaDefinition struct {
	Id       IdentifierExpression
	Fields   []*FieldExpression
	Generics []*TypeExpression
}

func (s *SchemaDefinition) ContainsNullableFields() bool {
	for _, field := range s.Fields {
		if field.IsOptional {
			return true
		}
	}

	return false
}

type AnnotationExpression struct {
	Type  syntax.AnnotationType
	Value string
}

type FieldExpression struct {
	Id          IdentifierExpression
	IsOptional  bool
	Type        *TypeExpression
	Annotations []*AnnotationExpression
}

type TypeExpression struct {
	Id       IdentifierExpression
	Generics []*TypeExpression
}

type EnumDefinition struct {
	Id     IdentifierExpression
	Values []string
}

type EskemaTree struct {
	Expr []*EskemaExpression
}
